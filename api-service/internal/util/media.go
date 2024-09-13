package util

import "io"

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

// Implement getImageDimensions function to get image width and height
func GetImageDimensions(file io.Reader) (int, int) {
	// You can use an image processing library like github.com/nfnt/resize
	// to get the dimensions of the image
	// For now, we'll return placeholder values
	return 800, 600
}
