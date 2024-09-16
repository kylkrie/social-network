package entity

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"yabro.io/social-api/internal/logger"
)

func CreateEntity[T any](ctx context.Context, exec sqlx.ExtContext, tableName string, entity T) error {
	// Get the type of the entity
	t := reflect.TypeOf(entity)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Build the column names and placeholders
	var columns []string
	var placeholders []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag != "" && dbTag != "-" {
			columns = append(columns, dbTag)
			placeholders = append(placeholders, ":"+dbTag)
		}
	}

	// Construct the query
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "))

	// Execute the query
	_, err := sqlx.NamedExecContext(ctx, exec, query, entity)
	if err != nil {
		return fmt.Errorf("failed to execute create query for %s: %w", tableName, err)
	}

	return nil
}

func CreateEntities[T any](ctx context.Context, exec sqlx.ExtContext, tableName string, entities []T) error {
	if len(entities) == 0 {
		return nil // Nothing to insert
	}

	// Reflect on the first entity to get structure information
	entityType := reflect.TypeOf(entities[0])
	if entityType.Kind() == reflect.Ptr {
		entityType = entityType.Elem()
	}

	var columns []string
	var placeholders []string

	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag != "" && dbTag != "-" {
			columns = append(columns, dbTag)
			placeholders = append(placeholders, ":"+dbTag)
		}
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	// Use type assertion to call NamedExecContext
	var err error
	switch e := exec.(type) {
	case *sqlx.DB:
		_, err = e.NamedExecContext(ctx, query, entities)
	case *sqlx.Tx:
		_, err = e.NamedExecContext(ctx, query, entities)
	default:
		return fmt.Errorf("unsupported exec type: %T", exec)
	}

	if err != nil {
		return fmt.Errorf("failed to bulk insert entities into %s: %w", tableName, err)
	}

	return nil
}

func GetEntity[T any](ctx context.Context, q sqlx.Queryer, ID int64, tableName string, idColName string) (*T, error) {
	var entity T
	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE %s = $1
	`, tableName, idColName)

	err := sqlx.Get(q, &entity, query, ID)
	if err != nil {
		logger.Warn(ctx).Str("table", tableName).Int64("id", ID).Err(err).Msg("GetEntity query failed")
		return nil, err
	}

	return &entity, nil
}

func GetEntities[T any](ctx context.Context, q sqlx.Queryer, IDs []int64, tableName string, idColName string) ([]T, error) {
	var entities []T
	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE %s = ANY($1)
		ORDER BY %s
	`, tableName, idColName, idColName)

	err := sqlx.Select(q, &entities, query, IDs)
	if err != nil {
		logger.Warn(ctx).Str("table", tableName).Err(err).Msg("GetEntities query failed")
		return nil, err
	}

	return entities, nil
}

func ListEntities[T any](ctx context.Context, q sqlx.Queryer, tableName string, idColName string, limit int, cursor *int64, where *string) ([]T, *int64, error) {
	var entities []T
	var query string
	var args []interface{}

	// Base query
	query = fmt.Sprintf("SELECT * FROM %s", tableName)

	// Add where condition if provided
	if where != nil {
		query += fmt.Sprintf(" WHERE %s", *where)
	}

	// Add cursor condition
	if cursor != nil {
		if where != nil {
			query += fmt.Sprintf(" AND %s < $1", idColName)
		} else {
			query += fmt.Sprintf(" WHERE %s < $1", idColName)
		}
		args = append(args, *cursor)
	}

	// Add ordering
	query += fmt.Sprintf(" ORDER BY %s DESC", idColName)

	// Add limit
	// Fetch one extra to determine if there are more results
	query += fmt.Sprintf(" LIMIT %d", limit+1)

	// Execute query
	err := sqlx.Select(q, &entities, query, args...)
	if err != nil {
		logger.Warn(ctx).Str("table", tableName).Err(err).Msg("ListEntities query failed")
		return nil, nil, err
	}

	// Determine if there are more results and set the next cursor
	var nextCursor *int64
	if len(entities) > limit {
		nextCursor = new(int64)

		// Get the ID of the last entity (before removing the extra one)
		lastEntity := reflect.ValueOf(entities[limit-1])
		t := reflect.TypeOf(entities[0])

		var idField reflect.Value
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.Tag.Get("db") == idColName {
				idField = lastEntity.Field(i)
				break
			}
		}

		if idField.IsValid() && idField.Kind() == reflect.Int64 {
			*nextCursor = idField.Int()
		} else {
			logger.Error(ctx).Str("table", tableName).Msg("Unable to set next cursor: ID field not found or invalid")
			nextCursor = nil
		}

		// Remove the extra entity
		entities = entities[:limit]
	}

	return entities, nextCursor, nil
}

func UpdateEntity[T any](ctx context.Context, db sqlx.ExtContext, tableName string, entity T, idColName string, idValue int64) error {
	v := reflect.ValueOf(entity)
	t := v.Type()

	var setClauses []string
	updateMap := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		dbTag := field.Tag.Get("db")
		if dbTag != "" && dbTag != "-" {
			// For non-pointer fields
			if field.Type.Kind() != reflect.Ptr {
				setClauses = append(setClauses, fmt.Sprintf("%s = :%s", dbTag, dbTag))
				updateMap[dbTag] = value.Interface()
			} else if !value.IsNil() { // For pointer fields that are not nil
				setClauses = append(setClauses, fmt.Sprintf("%s = :%s", dbTag, dbTag))
				updateMap[dbTag] = value.Elem().Interface()
			}
		}
	}

	if len(setClauses) == 0 {
		return nil // No fields to update
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = :id",
		tableName,
		strings.Join(setClauses, ", "),
		idColName,
	)

	// Add the ID to the update map
	updateMap["id"] = idValue

	_, err := sqlx.NamedExecContext(ctx, db, query, updateMap)
	if err != nil {
		logger.Warn(ctx).Str("table", tableName).Err(err).Msg("UpdateEntity query failed")
		return err
	}

	return nil
}
