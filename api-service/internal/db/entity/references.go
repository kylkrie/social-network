package entity

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"yabro.io/social-api/internal/logger"
	"yabro.io/social-api/internal/util"
)

type ReferenceEntity interface {
	GetReferenceID() int64
}

func GetEntitiesForID[T ReferenceEntity](ctx context.Context, q sqlx.Queryer, ID int64, tableName, refIdColName string) ([]T, error) {
	var entities []T
	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE %s = $1
		ORDER BY %s

	`, tableName, refIdColName, refIdColName)

	err := sqlx.Select(q, &entities, query, ID)
	if err != nil {
		logger.Warn(ctx).Str("table", tableName).Err(err).Msg("GetReferenceEntities query failed")
		return nil, err
	}

	return entities, nil
}

func GetEntitiesForIDs[T ReferenceEntity](ctx context.Context, q sqlx.Queryer, IDs []int64, tableName, refIdColName string, idColName *string) (map[int64][]T, error) {
	var entities []T
	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE %s = ANY($1)
		ORDER BY %s

	`, tableName, refIdColName, refIdColName)

	if idColName != nil {
		query += fmt.Sprintf(`, %s`, *idColName)
	}

	err := sqlx.Select(q, &entities, query, IDs)
	if err != nil {
		logger.Warn(ctx).Str("table", tableName).Err(err).Msg("GetReferenceEntities query failed")
		return nil, err
	}

	entityMap := util.ArrToMapArr(entities, func(e T) int64 { return e.GetReferenceID() })
	return entityMap, nil
}
