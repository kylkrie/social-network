package postdb

import (
	"fmt"
)

type CreatePostParams struct {
	ID             int64
	Content        string
	AuthorID       int64
	ConversationID *int64
	References     *[]CreatePostReference
}

type CreatePostReference struct {
	ID              int64
	ReferencePostID int64
	ReferenceType   PostReferenceType
}

func (pdb *PostDB) CreatePost(p CreatePostParams) error {
	tx, err := pdb.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Insert into posts table
	query := `
		INSERT INTO posts (id, content, author_id, conversation_id)
		VALUES ($1, $2, $3, $4)
	`
	_, err = tx.Exec(query, p.ID, p.Content, p.AuthorID, p.ConversationID)
	if err != nil {
		return fmt.Errorf("failed to create post: %w", err)
	}

	// Insert into post_references table and update metrics if references are provided
	if p.References != nil {
		references := *p.References
		if len(references) > 0 {
			referenceQuery := `
			INSERT INTO post_references (id, source_post_id, referenced_post_id, reference_type)
			VALUES ($1, $2, $3, $4)
			`
			for _, ref := range references {
				_, err := tx.Exec(referenceQuery, ref.ID, p.ID, ref.ReferencePostID, ref.ReferenceType)
				if err != nil {
					return fmt.Errorf("failed to create post reference: %w", err)
				}

				// If it's a reply, update the reply count for the referenced post
				if ref.ReferenceType == PostReferenceTypeReplyTo {
					updateMetricsQuery := `
					INSERT INTO post_public_metrics (post_id, replies)
					VALUES ($1, 1)
					ON CONFLICT (post_id) 
					DO UPDATE SET replies = post_public_metrics.replies + 1
					`
					_, err := tx.Exec(updateMetricsQuery, ref.ReferencePostID)
					if err != nil {
						return fmt.Errorf("failed to update reply metrics: %w", err)
					}
				}
			}
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
