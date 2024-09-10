package service

import "fmt"

func (s *PostService) BookmarkPost(postID, userID int64) error {
	err := s.postDB.BookmarkPost(postID, userID)
	if err != nil {
		return fmt.Errorf("failed to bookmark post: %w", err)
	}

	return nil
}

func (s *PostService) UnbookmarkPost(postID, userID int64) error {
	err := s.postDB.UnbookmarkPost(postID, userID)
	if err != nil {
		return fmt.Errorf("failed to unbookmark post: %w", err)
	}
	return nil
}
