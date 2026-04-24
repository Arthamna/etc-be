package repositories

import (
	"arthamna/rplLibrary/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *gorm.DB, category *models.Category) (*models.Category, error)
	FindAll(ctx context.Context) ([]models.Category, error)
	FindByID(ctx context.Context, id string) (*models.Category, error)
	Update(ctx context.Context, tx *gorm.DB, category *models.Category) (*models.Category, error)
	Delete(ctx context.Context, tx *gorm.DB, id string) error
	FindByIDs(ctx context.Context, ids []string) ([]models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) dbOrTx(tx *gorm.DB) *gorm.DB {
	if tx == nil {
		return r.db
	}
	return tx
}

func (r *categoryRepository) Create(ctx context.Context, tx *gorm.DB, category *models.Category) (*models.Category, error) {
	db := r.dbOrTx(tx)
	if err := db.WithContext(ctx).Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) FindAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) FindByID(ctx context.Context, id string) (*models.Category, error) {
	var category models.Category
	result := r.db.WithContext(ctx).First(&category, "category_id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, result.Error
	}
	return &category, nil
}

func (r *categoryRepository) Update(ctx context.Context, tx *gorm.DB, category *models.Category) (*models.Category, error) {
	db := r.dbOrTx(tx)
	if err := db.WithContext(ctx).Save(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) Delete(ctx context.Context, tx *gorm.DB, id string) error {
	db := r.dbOrTx(tx)
	return db.WithContext(ctx).Delete(&models.Category{}, "category_id = ?", id).Error
}

func (r *categoryRepository) FindByIDs(ctx context.Context, ids []string) ([]models.Category, error) {
	var categories []models.Category
	if len(ids) == 0 {
		return categories, nil
	}

	err := r.db.WithContext(ctx).
		Where("category_id IN ? AND deleted_at IS NULL", ids).
		Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}