package db

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"yabro.io/social-api/internal/util"
)

const (
	ProfileBucket = "profiles"
	MediaBucket   = "media"
)

type MinioStorage struct {
	client     *minio.Client
	cdnBaseURL string
}

func NewMinioStorage(client *minio.Client, cdnBaseURL string) *MinioStorage {
	return &MinioStorage{client: client, cdnBaseURL: cdnBaseURL}
}

func (m *MinioStorage) UploadFile(bucketName, objectName string, reader io.Reader, size int64) error {
	ctx := context.Background()

	// Upload the file
	_, err := m.client.PutObject(ctx, bucketName, objectName, reader, size, minio.PutObjectOptions{
		ContentType: util.GetContentType(objectName),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	return nil
}

func (m *MinioStorage) DeleteFile(bucketName, objectName string) error {
	ctx := context.Background()

	err := m.client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

func (m *MinioStorage) GetFileURL(bucketName, objectName string) (string, error) {
	// This is a simple implementation. You might want to use pre-signed URLs for more security
	return fmt.Sprintf("%s/%s/%s", m.cdnBaseURL, bucketName, objectName), nil
}
