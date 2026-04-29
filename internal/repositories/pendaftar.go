package repositories

import (
    "context"
    "etc-backend/internal/models"
    "gorm.io/gorm"
)

type PendaftarRepository interface {
    Create(ctx context.Context, pendaftar *models.Pendaftar) (*models.Pendaftar, error)
    FindByID(ctx context.Context, id string) (*models.Pendaftar, error)
    FindByUserAndRekrutmen(ctx context.Context, userID, rekrutmenID string) (*models.Pendaftar, error)
    UpdateStatus(ctx context.Context, pendaftarID, status string) error
}

type pendaftarRepository struct {
    db *gorm.DB
}

func NewPendaftarRepository(db *gorm.DB) PendaftarRepository {
    return &pendaftarRepository{db: db}
}

func (r *pendaftarRepository) Create(ctx context.Context, pendaftar *models.Pendaftar) (*models.Pendaftar, error) {
    if err := r.db.WithContext(ctx).Create(pendaftar).Error; err != nil {
        return nil, err
    }
    return pendaftar, nil
}

func (r *pendaftarRepository) FindByID(ctx context.Context, id string) (*models.Pendaftar, error) {
    var pendaftar models.Pendaftar
    if err := r.db.WithContext(ctx).
        Preload("User").
        Preload("Rekrutmen").
        Preload("Rekrutmen.User").
        First(&pendaftar, "pendaftar_id = ?", id).Error; err != nil {
        return nil, err
    }
    return &pendaftar, nil
}

func (r *pendaftarRepository) FindByUserAndRekrutmen(ctx context.Context, userID, rekrutmenID string) (*models.Pendaftar, error) {
    var pendaftar models.Pendaftar
    if err := r.db.WithContext(ctx).
        Where("user_id = ? AND rekrutmen_id = ?", userID, rekrutmenID).
        First(&pendaftar).Error; err != nil {
        return nil, err
    }
    return &pendaftar, nil
}

func (r *pendaftarRepository) UpdateStatus(ctx context.Context, pendaftarID, status string) error {
    return r.db.WithContext(ctx).
        Model(&models.Pendaftar{}).
        Where("pendaftar_id = ?", pendaftarID).
        Update("status", status).Error
}
