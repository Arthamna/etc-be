package repositories

import (
    "context"
    "etc-backend/internal/models"
    "gorm.io/gorm"
)

type HistoryRepository interface {
    Create(ctx context.Context, history *models.History) error
    FindByReviewerAndUserAndTim(ctx context.Context, reviewerID, userID, timID string) (*models.History, error)
    FindByUserID(ctx context.Context, userID string) ([]models.History, error)
}

type historyRepository struct {
    db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
    return &historyRepository{db: db}
}

func (r *historyRepository) Create(ctx context.Context, history *models.History) error {
    return r.db.WithContext(ctx).Create(history).Error
}

func (r *historyRepository) FindByReviewerAndUserAndTim(ctx context.Context, reviewerID, userID, timID string) (*models.History, error) {
    var history models.History
    if err := r.db.WithContext(ctx).
        Where("reviewer_user_id = ? AND user_id = ? AND tim_id = ?", reviewerID, userID, timID).
        First(&history).Error; err != nil {
        return nil, err
    }
    return &history, nil
}

func (r *historyRepository) FindByUserID(ctx context.Context, userID string) ([]models.History, error) {
	var histories []models.History
	err := r.db.WithContext(ctx).
		Preload("Reviewer").
		Preload("Tim").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&histories).Error
	return histories, err
}
