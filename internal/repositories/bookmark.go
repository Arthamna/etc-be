package repositories

import (
	"context"
	"etc-backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookmarkRepository interface {
	Add(ctx context.Context, userID, rekrutmenID string) error
	Delete(ctx context.Context, userID, rekrutmenID string) error
	FindByUser(ctx context.Context, userID string) ([]models.Bookmark, error)
}

type bookmarkRepository struct {
	db *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) BookmarkRepository {
	return &bookmarkRepository{db: db}
}

func (r *bookmarkRepository) Add(ctx context.Context, userID, rekrutmenID string) error {
	bm := models.Bookmark{
		ID:          uuid.NewString(),
		UserID:      userID,
		RekrutmenID: rekrutmenID,
	}

	return r.db.WithContext(ctx).
		FirstOrCreate(&bm, models.Bookmark{
			UserID:      userID,
			RekrutmenID: rekrutmenID,
		}).Error
}

func (r *bookmarkRepository) Delete(ctx context.Context, userID, rekrutmenID string) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND rekrutmen_id = ?", userID, rekrutmenID).
		Delete(&models.Bookmark{}).Error
}

func (r *bookmarkRepository) FindByUser(ctx context.Context, userID string) ([]models.Bookmark, error) {
	var bookmarks []models.Bookmark

	err := r.db.WithContext(ctx).
		Preload("Rekrutmen").
		Where("user_id = ?", userID).
		Find(&bookmarks).Error

	return bookmarks, err
}