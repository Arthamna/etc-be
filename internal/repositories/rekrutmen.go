package repositories

import (
	"context"
	"etc-backend/internal/models"

	"gorm.io/gorm"
)

type RekrutmenRepository interface {
	Create(ctx context.Context, rekrutmen *models.Rekrutmen) (*models.Rekrutmen, error)
	FindAll(ctx context.Context, page, limit int, kegiatan, role, keyword string) ([]models.Rekrutmen, int64, error)
	FindByID(ctx context.Context, id string) (*models.Rekrutmen, error)
	FindByUserID(ctx context.Context, userID string) ([]models.Rekrutmen, error)
	Update(ctx context.Context, rekrutmen *models.Rekrutmen) (*models.Rekrutmen, error)
	Delete(ctx context.Context, id string) error
	IsOwnedByUser(ctx context.Context, rekrutmenID, userID string) (bool, error)
}

type rekrutmenRepository struct {
	db *gorm.DB
}

func NewRekrutmenRepository(db *gorm.DB) RekrutmenRepository {
	return &rekrutmenRepository{db: db}
}

func (r *rekrutmenRepository) Create(ctx context.Context, rekrutmen *models.Rekrutmen) (*models.Rekrutmen, error) {
	if err := r.db.WithContext(ctx).Create(rekrutmen).Error; err != nil {
		return nil, err
	}
	return rekrutmen, nil
}

func (r *rekrutmenRepository) FindAll(ctx context.Context, page, limit int, kegiatan, role, keyword string) ([]models.Rekrutmen, int64, error) {
	var rekrutmen []models.Rekrutmen
	var total int64

	query := r.db.WithContext(ctx).Model(&models.Rekrutmen{}).Preload("User")

	if kegiatan != "" {
		query = query.Where("kegiatan = ?", kegiatan)
	}
	if role != "" {
		query = query.Where("role ILIKE ?", "%"+role+"%")
	}
	if keyword != "" {
		query = query.Where("role ILIKE ? OR kegiatan ILIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	query.Count(&total)

	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&rekrutmen).Error; err != nil {
		return nil, 0, err
	}

	return rekrutmen, total, nil
}

func (r *rekrutmenRepository) IsOwnedByUser(ctx context.Context, rekrutmenID, userID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.Rekrutmen{}).
		Where("rekrutmen_id = ? AND user_id = ?", rekrutmenID, userID).
		Count(&count).Error

	return count > 0, err
}

func (r *rekrutmenRepository) FindByID(ctx context.Context, id string) (*models.Rekrutmen, error) {
	var rekrutmen models.Rekrutmen
	if err := r.db.WithContext(ctx).Preload("User").First(&rekrutmen, "rekrutmen_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &rekrutmen, nil
}

func (r *rekrutmenRepository) FindByUserID(ctx context.Context, userID string) ([]models.Rekrutmen, error) {
	var rekrutmen []models.Rekrutmen
	if err := r.db.WithContext(ctx).Preload("User").Where("user_id = ?", userID).Find(&rekrutmen).Error; err != nil {
		return nil, err
	}
	return rekrutmen, nil
}

func (r *rekrutmenRepository) Update(ctx context.Context, rekrutmen *models.Rekrutmen) (*models.Rekrutmen, error) {
	if err := r.db.WithContext(ctx).Save(rekrutmen).Error; err != nil {
		return nil, err
	}
	return rekrutmen, nil
}

func (r *rekrutmenRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.Rekrutmen{}, "rekrutmen_id = ?", id).Error
}
