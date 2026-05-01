package repositories

import (
    "context"
    "etc-backend/internal/models"
    "gorm.io/gorm"
)

type TimRepository interface {
    Create(ctx context.Context, tim *models.Tim) error
    FindByID(ctx context.Context, timID string) (*models.Tim, error)
    FindByRekrutmenID(ctx context.Context, rekrutmenID string) (*models.Tim, error)
    AddParticipant(ctx context.Context, participant *models.TimParticipant) error
    CountParticipants(ctx context.Context, timID string) (int64, error)
    IsParticipant(ctx context.Context, timID, userID string) (bool, error)
    RemoveParticipant(ctx context.Context, timID, userID string) error
}

type timRepository struct {
    db *gorm.DB
}

func NewTimRepository(db *gorm.DB) TimRepository {
    return &timRepository{db: db}
}

func (r *timRepository) Create(ctx context.Context, tim *models.Tim) error {
    return r.db.WithContext(ctx).Create(tim).Error
}

func (r *timRepository) FindByID(ctx context.Context, timID string) (*models.Tim, error) {
    var tim models.Tim
    if err := r.db.WithContext(ctx).
        Preload("Participants").
        Preload("Participants.User").
        First(&tim, "tim_id = ?", timID).Error; err != nil {
        return nil, err
    }
    return &tim, nil
}

func (r *timRepository) FindByRekrutmenID(ctx context.Context, rekrutmenID string) (*models.Tim, error) {
    var tim models.Tim
    if err := r.db.WithContext(ctx).
        Where("rekrutmen_id = ?", rekrutmenID).
        First(&tim).Error; err != nil {
        return nil, err
    }
    return &tim, nil
}

func (r *timRepository) AddParticipant(ctx context.Context, participant *models.TimParticipant) error {
    return r.db.WithContext(ctx).Create(participant).Error
}

func (r *timRepository) CountParticipants(ctx context.Context, timID string) (int64, error) {
    var count int64
    err := r.db.WithContext(ctx).
        Model(&models.TimParticipant{}).
        Where("tim_id = ?", timID).
        Count(&count).Error
    return count, err
}

func (r *timRepository) IsParticipant(ctx context.Context, timID, userID string) (bool, error) {
    var count int64
    err := r.db.WithContext(ctx).
        Model(&models.TimParticipant{}).
        Where("tim_id = ? AND user_id = ?", timID, userID).
        Count(&count).Error
    return count > 0, err
}

func (r *timRepository) RemoveParticipant(ctx context.Context, timID, userID string) error {
    return r.db.WithContext(ctx).
        Where("tim_id = ? AND user_id = ?", timID, userID).
        Delete(&models.TimParticipant{}).Error
}
