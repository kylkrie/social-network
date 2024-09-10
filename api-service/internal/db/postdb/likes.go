package postdb

import (
	"fmt"
)

func (pdb *PostDB) LikePost(postID, userID int64) error {
	tx, err := pdb.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
        INSERT INTO post_likes (post_id, user_id)
        VALUES ($1, $2)
        ON CONFLICT (post_id, user_id) DO NOTHING
    `
	result, err := tx.Exec(query, postID, userID)
	if err != nil {
		return fmt.Errorf("failed to like post: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected > 0 {
		// Only update metrics if a new like was actually inserted
		updateQuery := `
            INSERT INTO post_public_metrics (post_id, likes)
            VALUES ($1, 1)
            ON CONFLICT (post_id)
            DO UPDATE SET likes = post_public_metrics.likes + 1
        `
		_, err = tx.Exec(updateQuery, postID)
		if err != nil {
			return fmt.Errorf("failed to update post metrics: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (pdb *PostDB) UnlikePost(postID, userID int64) error {
	tx, err := pdb.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
        DELETE FROM post_likes
        WHERE post_id = $1 AND user_id = $2
    `
	result, err := tx.Exec(query, postID, userID)
	if err != nil {
		return fmt.Errorf("failed to unlike post: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected > 0 {
		// Only update metrics if a like was actually removed
		updateQuery := `
            UPDATE post_public_metrics
            SET likes = GREATEST(likes - 1, 0)
            WHERE post_id = $1
        `
		_, err = tx.Exec(updateQuery, postID)
		if err != nil {
			return fmt.Errorf("failed to update post metrics: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

type UserPostInteraction struct {
	PostID       int64
	IsLiked      bool
	IsBookmarked bool
}

func (pdb *PostDB) GetUserPostInteractions(postIDs []int64, userID int64) ([]UserPostInteraction, error) {
	query := `
        SELECT 
            p.id AS post_id,
            CASE WHEN l.user_id IS NOT NULL THEN true ELSE false END AS is_liked,
            CASE WHEN b.user_id IS NOT NULL THEN true ELSE false END AS is_bookmarked
        FROM unnest($1::bigint[]) p(id)
        LEFT JOIN post_likes l ON p.id = l.post_id AND l.user_id = $2
        LEFT JOIN post_bookmarks b ON p.id = b.post_id AND b.user_id = $2
        ORDER BY p.id
    `

	rows, err := pdb.db.Query(query, postIDs, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user post interactions: %w", err)
	}
	defer rows.Close()

	var interactions []UserPostInteraction
	for rows.Next() {
		var interaction UserPostInteraction
		err := rows.Scan(&interaction.PostID, &interaction.IsLiked, &interaction.IsBookmarked)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user post interaction: %w", err)
		}
		interactions = append(interactions, interaction)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user post interactions: %w", err)
	}

	return interactions, nil
}
