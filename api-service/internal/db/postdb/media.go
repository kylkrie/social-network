package postdb

import (
	"fmt"
)

func (pdb *PostDB) AddMediaToPost(postID int64, media []PostMedia) error {
	query := `
		INSERT INTO post_media (media_key, post_id, type, url, width, height)
		VALUES (:media_key, :post_id, :type, :url, :width, :height)
	`

	for i := range media {
		media[i].PostID = postID
	}

	_, err := pdb.db.NamedExec(query, media)
	if err != nil {
		return fmt.Errorf("failed to add media to post: %w", err)
	}

	return nil
}

func (pdb *PostDB) GetMediaForPost(postID int64) ([]PostMedia, error) {
	var media []PostMedia
	query := `
		SELECT * FROM post_media
		WHERE post_id = $1
		ORDER BY media_key
	`

	err := pdb.db.Select(&media, query, postID)
	if err != nil {
		return nil, fmt.Errorf("failed to get media for post: %w", err)
	}

	return media, nil
}

func (pdb *PostDB) GetMediaForPosts(postIDs []int64) (map[int64][]PostMedia, error) {
	var media []PostMedia
	query := `
		SELECT * FROM post_media
		WHERE post_id = ANY($1)
		ORDER BY post_id, media_key
	`

	err := pdb.db.Select(&media, query, postIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get media for posts: %w", err)
	}

	mediaMap := make(map[int64][]PostMedia)
	for _, m := range media {
		mediaMap[m.PostID] = append(mediaMap[m.PostID], m)
	}

	return mediaMap, nil
}
