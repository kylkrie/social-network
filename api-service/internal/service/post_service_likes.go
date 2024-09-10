package service

import "fmt"

func (s *PostService) LikePost(postID, userID int64) error {
	err := s.postDB.LikePost(postID, userID)
	if err != nil {
		return fmt.Errorf("failed to like post: %w", err)
	}
	return nil
}

func (s *PostService) UnlikePost(postID, userID int64) error {
	err := s.postDB.UnlikePost(postID, userID)
	if err != nil {
		return fmt.Errorf("failed to unlike post: %w", err)
	}
	return nil
}
