package repositories

import (
	"arthamna/rplLibrary/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(ctx context.Context, tx *gorm.DB, book *models.Book) (*models.Book, error)
	FindAll(ctx context.Context) ([]models.Book, error)
	FindByID(ctx context.Context, id string) (*models.Book, error)
	FindByTitle(ctx context.Context, query string) ([]models.Book, error)
	Update(ctx context.Context, tx *gorm.DB, book *models.Book) (*models.Book, error)
	UpdatePicture(ctx context.Context, tx *gorm.DB, book *models.Book) (*models.Book, error)
	Delete(ctx context.Context, tx *gorm.DB, id string) error
	FindByCategory(ctx context.Context, categoryName string) ([]models.Book, error)
	FindByStatus(ctx context.Context, status string) ([]models.Book, error)
	AssignCategories(ctx context.Context, tx *gorm.DB, bookID string, categories []models.Category) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) dbOrTx(tx *gorm.DB) *gorm.DB {
	if tx == nil {
		return r.db
	}
	return tx
}

func (r *bookRepository) preloadBookRelations(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Categories").
		Preload("Borrowings")
}

func (r *bookRepository) Create(ctx context.Context, tx *gorm.DB, book *models.Book) (*models.Book, error) {
	db := r.dbOrTx(tx)
	if err := db.WithContext(ctx).Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *bookRepository) FindAll(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	err := r.preloadBookRelations(r.db.WithContext(ctx)).Find(&books).Error
	return books, err
}

func (r *bookRepository) FindByID(ctx context.Context, id string) (*models.Book, error) {
	var book models.Book
	result := r.preloadBookRelations(r.db.WithContext(ctx)).First(&book, "book_id = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, result.Error
	}

	return &book, nil
}

func (r *bookRepository) FindByTitle(ctx context.Context, query string) ([]models.Book, error) {
    var books []models.Book
    err := r.db.WithContext(ctx).
        Preload("Categories").
        Where("LOWER(title) LIKE LOWER(?)", "%"+query+"%").
        Find(&books).Error
    if err != nil {
        return nil, err
    }
    return books, nil
}

func (r *bookRepository) Update(ctx context.Context, tx *gorm.DB, book *models.Book) (*models.Book, error) {
	db := r.dbOrTx(tx)
	if err := db.WithContext(ctx).Save(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *bookRepository) UpdatePicture(ctx context.Context, tx *gorm.DB, book *models.Book) (*models.Book, error) {
	db := r.dbOrTx(tx)
	if err := db.WithContext(ctx).	
	Model(&models.Book{}).
	Where("book_id = ?", book.BookID).
	Update("book_picture", book.BookPicture).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *bookRepository) Delete(ctx context.Context, tx *gorm.DB, id string) error {
	db := r.dbOrTx(tx)
	return db.WithContext(ctx).Delete(&models.Book{}, "book_id = ?", id).Error
}

func (r *bookRepository) FindByCategory(ctx context.Context, categoryName string) ([]models.Book, error) {
	var books []models.Book
	err := r.preloadBookRelations(r.db.WithContext(ctx)).
		Model(&models.Book{}).
		Distinct("books.*").
		Joins("JOIN book_categories bc ON bc.book_id = books.book_id").
		Joins("JOIN categories c ON c.category_id = bc.category_id").
		Where("LOWER(c.name) = LOWER(?) AND c.deleted_at IS NULL", categoryName).
		Find(&books).Error

	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) FindByStatus(ctx context.Context, status string) ([]models.Book, error) {
	var books []models.Book
	err := r.preloadBookRelations(r.db.WithContext(ctx)).
		Where("status = ?", status).
		Find(&books).Error

	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) AssignCategories(ctx context.Context, tx *gorm.DB, bookID string, categories []models.Category) error {
	db := r.dbOrTx(tx)
	book := models.Book{BookID: bookID}
	return db.WithContext(ctx).Model(&book).Association("Categories").Replace(&categories)
}