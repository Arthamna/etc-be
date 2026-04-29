package repositories

import (
	"etc-backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, user *models.User) (*models.User, error)
	FindByNoPengenal(ctx context.Context, nid string) (*models.User, error)
	FindAll(ctx context.Context) ([]models.User, error)
	FindByID(ctx context.Context,  id string) (*models.User, error)
	Update(ctx context.Context, tx *gorm.DB, user *models.User) (*models.User, error)
	Delete(ctx context.Context, tx *gorm.DB, id string) error
	FindBookmarksByUserID(ctx context.Context, userID string) ([]models.Bookmark, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) dbOrTx(tx *gorm.DB) *gorm.DB {
	if tx == nil {
		return r.db
	}
	return tx
}

func (r *userRepository) Create(ctx context.Context, tx *gorm.DB, user *models.User) (*models.User, error) {
	db := r.dbOrTx(tx)
	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByNoPengenal(ctx context.Context, nid string) (*models.User, error){
	var user models.User
	result := r.db.WithContext(ctx).Where("no_pengenal = ?", nid).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := r.db.WithContext(ctx).Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).First(&user, "user_id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, tx *gorm.DB, user *models.User) (*models.User, error) {
	db := r.dbOrTx(tx)
	if err := db.WithContext(ctx).Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, tx *gorm.DB, id string) error {
	db := r.dbOrTx(tx)
	return db.WithContext(ctx).Delete(&models.User{}, "user_id = ?", id).Error
}

func (r *userRepository) FindBookmarksByUserID(ctx context.Context, userID string) ([]models.Bookmark, error) {
    var bookmarks []models.Bookmark
    if err := r.db.WithContext(ctx).
        Preload("Rekrutmen").
        Preload("Rekrutmen.User").
        Where("user_id = ?", userID).
        Find(&bookmarks).Error; err != nil {
        return nil, err
    }
    return bookmarks, nil
}
