package repositories

import (
	"arthamna/rplLibrary/internal/models"
	"context"

	"gorm.io/gorm"
)

type BookBorrowingRepository interface {
	Create(ctx context.Context, tx *gorm.DB, bookBorrowing *models.BookBorrowing) (*models.BookBorrowing, error)
	FindCurrentlyByBookID(ctx context.Context, bookID string) (*models.BookBorrowing, error)
	Update(ctx context.Context, tx *gorm.DB, bookBorrowing *models.BookBorrowing) (*models.BookBorrowing, error)
}

type bookBorrowingRepository struct {
	db *gorm.DB
}

func NewBookBorrowingRepository(db *gorm.DB) BookBorrowingRepository {
	return &bookBorrowingRepository{db: db}
}

func (r *bookBorrowingRepository) dbOrTx(tx *gorm.DB) *gorm.DB {
	if tx == nil {
		return r.db
	}
	return tx
}

func (r *bookBorrowingRepository) Create(ctx context.Context, tx *gorm.DB, bookBorrowing *models.BookBorrowing) (*models.BookBorrowing, error) {
	db := r.dbOrTx(tx)
	if err := db.WithContext(ctx).Create(bookBorrowing).Error; err != nil {
		return nil, err
	}
	return bookBorrowing, nil
}

func (r *bookBorrowingRepository) FindCurrentlyByBookID(ctx context.Context, bookID string) (*models.BookBorrowing, error) {
	var bookBorrowing models.BookBorrowing
	err := r.db.WithContext(ctx).
		Where("book_id = ? AND returned_at IS NULL", bookID).
		First(&bookBorrowing).Error
	if err != nil {
		return nil, err
	}

	return &bookBorrowing, nil
}

func (r *bookBorrowingRepository) Update(ctx context.Context, tx *gorm.DB, bookBorrowing *models.BookBorrowing) (*models.BookBorrowing, error) {
	db := r.dbOrTx(tx)
	if err := db.WithContext(ctx).Save(bookBorrowing).Error; err != nil {
		return nil, err
	}
	return bookBorrowing, nil
}