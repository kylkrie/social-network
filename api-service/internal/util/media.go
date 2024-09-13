package util

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"path/filepath"
)

func GetContentType(filename string) string {
	ext := filepath.Ext(filename)
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	default:
		return "application/octet-stream"
	}
}

func GetMediaType(contentType string) string {
	switch contentType {
	case "image/jpeg", "image/png":
		return "photo"
	case "image/gif":
		return "animated_gif"
	case "video/mp4":
		return "video"
	default:
		return "unknown"
	}
}

func GetImageDimensions(file io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get image dimenstions: %v", err)
	}
	return img.Width, img.Height, nil
}
