package services

import (
	"context"
	"etc-backend/internal/models"
	"etc-backend/internal/repositories"
)

type (
	BookmarkService interface {
		AddBookmark(ctx context.Context, userID, rekrutmenID string) error 
		RemoveBookmark(ctx context.Context, userID, rekrutmenID string) error
		GetBookmarks(ctx context.Context, userID string) ([]models.Bookmark, error)
	}

	bookmarkService struct {
		bookmarkRepo repositories.BookmarkRepository
	}
)

func NewBookmarkService(bookmarkRepo repositories.BookmarkRepository) BookmarkService {
	return &bookmarkService{
		bookmarkRepo: bookmarkRepo,
	}
}

func (s *bookmarkService) AddBookmark(ctx context.Context, userID, rekrutmenID string) error {
	return s.bookmarkRepo.Add(ctx, userID, rekrutmenID)
}

func (s *bookmarkService) RemoveBookmark(ctx context.Context, userID, rekrutmenID string) error {
	return s.bookmarkRepo.Delete(ctx, userID, rekrutmenID)
}

func (s *bookmarkService) GetBookmarks(ctx context.Context, userID string) ([]models.Bookmark, error) {
	return s.bookmarkRepo.FindByUser(ctx, userID)
}