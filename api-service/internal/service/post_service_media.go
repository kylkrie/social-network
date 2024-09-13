package service

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"yabro.io/social-api/internal/db"
	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/util"
)

func (s *PostService) uploadMedia(postID int64, files []*multipart.FileHeader) ([]postdb.PostMedia, error) {
	var media []postdb.PostMedia

	for _, file := range files {
		mediaKey := s.snowflakeNode.Generate().Int64()
		ext := filepath.Ext(file.Filename)
		objectName := fmt.Sprintf("%d/%d%s", postID, mediaKey, ext)

		src, err := file.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
		defer src.Close()

		err = s.minioStorage.UploadFile(db.MediaBucket, objectName, src, file.Size)
		if err != nil {
			return nil, fmt.Errorf("failed to upload file: %w", err)
		}

		url, err := s.minioStorage.GetFileURL(db.MediaBucket, objectName)
		if err != nil {
			return nil, fmt.Errorf("failed to get file URL: %w", err)
		}

		mediaType := util.GetMediaType(file.Header.Get("Content-Type"))

		width, height, err := util.GetImageDimensions(src)
		if err != nil {
			return nil, err
		}

		media = append(media, postdb.PostMedia{
			MediaKey: mediaKey,
			PostID:   postID,
			Type:     mediaType,

			URL: url,

			Width: width,

			Height: height,
		})
	}

	return media, nil
}
