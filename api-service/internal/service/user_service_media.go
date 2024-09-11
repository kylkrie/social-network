package service

import (
	"fmt"
	"io"
	"path/filepath"

	"yabro.io/social-api/internal/db"
)

func (s *UserService) UploadProfilePicture(userID int64, reader io.Reader, size int64, filename string) error {
	// Generate a unique ID for the profile picture
	picID := s.snowflakeNode.Generate().Int64()

	// Extract the file extension
	ext := filepath.Ext(filename)
	objectName := fmt.Sprintf("%d/pfp/%d%s", userID, picID, ext)

	err := s.minioStorage.UploadFile(db.ProfileBucket, objectName, reader, size)
	if err != nil {
		return fmt.Errorf("failed to upload profile picture: %w", err)
	}

	// Get the URL of the uploaded file
	url, err := s.minioStorage.GetFileURL(db.ProfileBucket, objectName)
	if err != nil {
		return fmt.Errorf("failed to get profile picture URL: %w", err)
	}

	// Update the user's profile picture URL in the database
	err = s.userDB.UpdateProfilePictureURL(userID, url)
	if err != nil {
		return fmt.Errorf("failed to update profile picture URL: %w", err)
	}

	return nil
}

func (s *UserService) UploadProfileBanner(userID int64, reader io.Reader, size int64, filename string) error {
	// Generate a unique ID for the profile picture
	picID := s.snowflakeNode.Generate().Int64()

	// Extract the file extension
	ext := filepath.Ext(filename)
	objectName := fmt.Sprintf("%d/banner/%d%s", userID, picID, ext)

	err := s.minioStorage.UploadFile(db.ProfileBucket, objectName, reader, size)
	if err != nil {
		return fmt.Errorf("failed to upload profile banner: %w", err)
	}

	// Get the URL of the uploaded file
	url, err := s.minioStorage.GetFileURL(db.ProfileBucket, objectName)
	if err != nil {
		return fmt.Errorf("failed to get profile banner URL: %w", err)
	}

	// Update the user's profile picture URL in the database
	err = s.userDB.UpdateProfileBannerURL(userID, url)
	if err != nil {
		return fmt.Errorf("failed to update profile banner URL: %w", err)
	}

	return nil
}
