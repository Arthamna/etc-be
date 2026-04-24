package services

import (
	"arthamna/rplLibrary/constants"
	"arthamna/rplLibrary/internal/dtos"
	"arthamna/rplLibrary/internal/models"
	"arthamna/rplLibrary/internal/repositories"
	"context"
	"encoding/base64"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	BookService interface {
		Create(ctx context.Context, req dtos.BookCreateRequest) (dtos.BookResponse, error)
		UploadBookPicture(ctx context.Context, req dtos.UploadBookPictureRequest) (dtos.UploadBookPictureResponse, error)
		GetAll(ctx context.Context) ([]dtos.BookResponse, error)
		FindByStatus(ctx context.Context, status string) ([]dtos.BookResponse, error)
		FindByCategory(ctx context.Context, category string) ([]dtos.BookResponse, error)
		SearchByTitle(ctx context.Context, query string) ([]dtos.BookResponse, error)
		GetByID(ctx context.Context, bookID string) (dtos.BookResponse, error)
		Update(ctx context.Context, req dtos.BookUpdateRequest) (dtos.BookResponse, error)
		Delete(ctx context.Context, bookID string) error
		BorrowBook(ctx context.Context, req dtos.BorrowBookRequest, userId string) (dtos.BorrowBookResponse, error)
		BorrowMultipleBook(ctx context.Context, req dtos.BorrowMultipleBookRequest, userId string) (dtos.BorrowMultipleBookResponse, error)
		SetMultipleBookReturned(ctx context.Context, req dtos.SetMultipleReturnedRequest) (dtos.SetMultipleReturnedResponse, error)
		SetBookReturned(ctx context.Context, bookID string) (dtos.SetBookReturnedResponse, error)
	}

	bookService struct {
		bookRepo      repositories.BookRepository
		categoryRepo  repositories.CategoryRepository
		userRepo      repositories.UserRepository
		borrowingRepo repositories.BookBorrowingRepository
	}
)

func NewBookService(bookRepo repositories.BookRepository, categoryRepo repositories.CategoryRepository, userRepo repositories.UserRepository, borrowingRepo repositories.BookBorrowingRepository) BookService {
	return &bookService{
		bookRepo:      bookRepo,
		categoryRepo:  categoryRepo,
		userRepo:      userRepo,
		borrowingRepo: borrowingRepo,
	}
}

func (s *bookService) Create(ctx context.Context, req dtos.BookCreateRequest) (dtos.BookResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	book := &models.Book{
		BookID:      uuid.NewString(),
		Author:      req.Author,
		Title:       req.Title,
		Description: req.Description,
		Status:      constants.STATUS_AVAILABLE,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	createdBook, err := s.bookRepo.Create(ctx, nil, book)
	if err != nil {
		return dtos.BookResponse{}, err
	}

	categories, err := s.categoryRepo.FindByIDs(ctx, req.CategoryIDs)
	if err != nil {
		return dtos.BookResponse{}, err
	}

	if len(categories) != len(req.CategoryIDs) {
		return dtos.BookResponse{}, errors.New("one or more categories not found")
	}

	if err := s.bookRepo.AssignCategories(ctx, nil, book.BookID, categories); err != nil {
		return dtos.BookResponse{}, err
	}

	return dtos.BookResponse{
		BookID:      createdBook.BookID,
		Author:      createdBook.Author,
		Title:       createdBook.Title,
		Description: createdBook.Description,
		Status:      createdBook.Status,
		CategoryIDs: req.CategoryIDs,
		CreatedAt:   createdBook.CreatedAt,
		UpdatedAt:   createdBook.UpdatedAt,
	}, nil
}

func (s *bookService) UploadBookPicture(ctx context.Context, req dtos.UploadBookPictureRequest) (dtos.UploadBookPictureResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	if req.BookPicture == nil {
		return dtos.UploadBookPictureResponse{}, errors.New("book picture is required")
	}

	book, err := s.bookRepo.FindByID(ctx, req.BookID)
	if err != nil {
		return dtos.UploadBookPictureResponse{}, err
	}

	imageBytes, err := readImageBytes(req.BookPicture)
	if err != nil {
		return dtos.UploadBookPictureResponse{}, err
	}

	book.BookPicture = imageBytes
	book.UpdatedAt = time.Now()

	updatedBook, err := s.bookRepo.UpdatePicture(ctx, nil, book)
	if err != nil {
		return dtos.UploadBookPictureResponse{}, err
	}

	return dtos.UploadBookPictureResponse{
		BookID:      updatedBook.BookID,
		BookPicture: base64.StdEncoding.EncodeToString(updatedBook.BookPicture),
	}, nil
}

func (s *bookService) GetAll(ctx context.Context) ([]dtos.BookResponse, error) {
	books, err := s.bookRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var response []dtos.BookResponse
	for _, book := range books {
		var categoryIDs []string
		for _, category := range book.Categories {
			categoryIDs = append(categoryIDs, category.CategoryID)
		}

		response = append(response, dtos.BookResponse{
			BookID:      book.BookID,
			Title:       book.Title,
			Description: book.Description,
			Author:      book.Author,
			CategoryIDs: categoryIDs,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
		})
	}

	return response, nil
}

func (s *bookService) SearchByTitle(ctx context.Context, query string) ([]dtos.BookResponse, error) {
    if query == "" {
        return nil, errors.New("search query cannot be empty")
    }

    books, err := s.bookRepo.FindByTitle(ctx, query)
    if err != nil {
        return nil, err
    }

    var response []dtos.BookResponse
    for _, book := range books {
        book := book
        var categoryIDs []string
        for _, category := range book.Categories {
            categoryIDs = append(categoryIDs, category.CategoryID)
        }
        response = append(response, dtos.BookResponse{
            BookID:      book.BookID,
            Author:      book.Author,
            Title:       book.Title,
            Description: book.Description,
            Status:      book.Status,
            CategoryIDs: categoryIDs,
            CreatedAt:   book.CreatedAt,
            UpdatedAt:   book.UpdatedAt,
        })
    }

    return response, nil
}

func (s *bookService) FindByStatus(ctx context.Context, status string) ([]dtos.BookResponse, error) {
	books, err := s.bookRepo.FindByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	var response []dtos.BookResponse
	for _, book := range books {
		var categoryIDs []string
		for _, category := range book.Categories {
			categoryIDs = append(categoryIDs, category.CategoryID)
		}

		response = append(response, dtos.BookResponse{
			BookID:      book.BookID,
			Title:       book.Title,
			Description: book.Description,
			Author:      book.Author,
			CategoryIDs: categoryIDs,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
		})
	}

	return response, nil
}

func (s *bookService) FindByCategory(ctx context.Context, category string) ([]dtos.BookResponse, error) {
	books, err := s.bookRepo.FindByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	var response []dtos.BookResponse
	for _, book := range books {
		var categoryIDs []string
		for _, category := range book.Categories {
			categoryIDs = append(categoryIDs, category.CategoryID)
		}

		response = append(response, dtos.BookResponse{
			BookID:      book.BookID,
			Title:       book.Title,
			Description: book.Description,
			Author:      book.Author,
			CategoryIDs: categoryIDs,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
		})
	}

	return response, nil
}

func (s *bookService) GetByID(ctx context.Context, bookID string) (dtos.BookResponse, error) {
	book, err := s.bookRepo.FindByID(ctx, bookID)
	if err != nil {
		return dtos.BookResponse{}, err
	}

	var categoryIDs []string
	for _, category := range book.Categories {
		categoryIDs = append(categoryIDs, category.CategoryID)
	}

	return dtos.BookResponse{
		BookID:      book.BookID,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
		CategoryIDs: categoryIDs,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}, nil
}

func (s *bookService) Update(ctx context.Context, req dtos.BookUpdateRequest) (dtos.BookResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	book, err := s.bookRepo.FindByID(ctx, req.BookID)
	if err != nil {
		return dtos.BookResponse{}, err
	}

	if req.Author != "" {
		book.Author = req.Author
	}
	if req.Title != "" {
		book.Title = req.Title
	}
	if req.Description != "" {
		book.Description = req.Description
	}

	book.UpdatedAt = time.Now()

	updatedBook, err := s.bookRepo.Update(ctx, nil, book)
	if err != nil {
		return dtos.BookResponse{}, err
	}

	if len(req.CategoryIDs) > 0 {
		categories, err := s.categoryRepo.FindByIDs(ctx, req.CategoryIDs)
		if err != nil {
			return dtos.BookResponse{}, err
		}
		if len(categories) != len(req.CategoryIDs) {
			return dtos.BookResponse{}, errors.New("one or more categories not found")
		}

		if err := s.bookRepo.AssignCategories(ctx, nil, updatedBook.BookID, categories); err != nil {
			return dtos.BookResponse{}, err
		}
	}

	// reassign ke categories 
	updatedBook, err = s.bookRepo.FindByID(ctx, updatedBook.BookID)
	if err != nil {
		return dtos.BookResponse{}, err
	}

	var categoryIDs []string
	for _, c := range updatedBook.Categories {
		categoryIDs = append(categoryIDs, c.CategoryID)
	}

	return dtos.BookResponse{
		BookID:      updatedBook.BookID,
		Author:      updatedBook.Author,
		Title:       updatedBook.Title,
		Description: updatedBook.Description,
		Status:      updatedBook.Status,
		CategoryIDs: categoryIDs,
		CreatedAt:   updatedBook.CreatedAt,
		UpdatedAt:   updatedBook.UpdatedAt,
	}, nil
}

func (s *bookService) BorrowBook(ctx context.Context, req dtos.BorrowBookRequest, userId string) (dtos.BorrowBookResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	user, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		return dtos.BorrowBookResponse{}, err
	}

	book, err := s.bookRepo.FindByID(ctx, req.BookID)
	if err != nil {
		return dtos.BorrowBookResponse{}, err
	}

	if book.Status == constants.STATUS_BORROWED {
		return dtos.BorrowBookResponse{}, errors.New("book is already borrowed")
	}

	now := time.Now()
	borrowing := &models.BookBorrowing{
		BorrowingID: uuid.NewString(),
		UserID:      user.UserID,
		BookID:      book.BookID,
		BorrowedAt:  now,
	}

	createdBorrowing, err := s.borrowingRepo.Create(ctx, nil, borrowing)
	if err != nil {
		return dtos.BorrowBookResponse{}, err
	}

	book.Status = constants.STATUS_BORROWED
	book.UpdatedAt = now

	updatedBook, err := s.bookRepo.Update(ctx, nil, book)
	if err != nil {
		return dtos.BorrowBookResponse{}, err
	}

	return dtos.BorrowBookResponse{
		BookID:     updatedBook.BookID,
		Title:      updatedBook.Title,
		UserID:     user.UserID,
		Username:   user.Username,
		BorrowedAt: createdBorrowing.BorrowedAt,
	}, nil
}

func (s *bookService) BorrowMultipleBook(ctx context.Context, req dtos.BorrowMultipleBookRequest, userId string) (dtos.BorrowMultipleBookResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	user, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		return dtos.BorrowMultipleBookResponse{}, err
	}

	var results []dtos.BorrowBookResponse
	now := time.Now()

	for _, bookID := range req.BookIDs {
		book, err := s.bookRepo.FindByID(ctx, bookID)
		if err != nil {
			return dtos.BorrowMultipleBookResponse{}, errors.New("book not found: " + bookID)
		}
		if book.Status == constants.STATUS_BORROWED {
			return dtos.BorrowMultipleBookResponse{}, errors.New("book already borrowed: " + bookID)
		}

		borrowing := &models.BookBorrowing{
			BorrowingID: uuid.NewString(),
			UserID:      user.UserID,
			BookID:      book.BookID,
			BorrowedAt:  now,
		}
		createdBorrowing, err := s.borrowingRepo.Create(ctx, nil, borrowing)
		if err != nil {
			return dtos.BorrowMultipleBookResponse{}, err
		}

		book.Status = constants.STATUS_BORROWED
		book.UpdatedAt = now
		updatedBook, err := s.bookRepo.Update(ctx, nil, book)
		if err != nil {
			return dtos.BorrowMultipleBookResponse{}, err
		}

		results = append(results, dtos.BorrowBookResponse{
			BookID:     updatedBook.BookID,
			Title:      updatedBook.Title,
			UserID:     user.UserID,
			Username:   user.Username,
			BorrowedAt: createdBorrowing.BorrowedAt,
		})
	}

	return dtos.BorrowMultipleBookResponse{Borrowed: results}, nil
}


func (s *bookService) SetBookReturned(ctx context.Context, bookID string) (dtos.SetBookReturnedResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	book, err := s.bookRepo.FindByID(ctx, bookID)
	if err != nil {
		return dtos.SetBookReturnedResponse{}, err
	}

	if book.Status == constants.STATUS_AVAILABLE {
		return dtos.SetBookReturnedResponse{}, errors.New("book is already available")
	}

	borrowing, err := s.borrowingRepo.FindCurrentlyByBookID(ctx, bookID)
	if err != nil {
		return dtos.SetBookReturnedResponse{}, err
	}

	now := time.Now()
	borrowing.ReturnedAt = &now

	updatedBorrowing, err := s.borrowingRepo.Update(ctx, nil, borrowing)
	if err != nil {
		return dtos.SetBookReturnedResponse{}, err
	}

	book.Status = constants.STATUS_AVAILABLE
	book.UpdatedAt = now

	updatedBook, err := s.bookRepo.Update(ctx, nil, book)
	if err != nil {
		return dtos.SetBookReturnedResponse{}, err
	}

	return dtos.SetBookReturnedResponse{
		BookID:      updatedBook.BookID,
		BorrowingID: updatedBorrowing.BorrowingID,
		ReturnedAt:  *updatedBorrowing.ReturnedAt,
	}, nil
}

func (s *bookService) SetMultipleBookReturned(ctx context.Context, req dtos.SetMultipleReturnedRequest) (dtos.SetMultipleReturnedResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	var results []dtos.SetBookReturnedResponse
	now := time.Now()

	for _, bookID := range req.BookIDs {
		book, err := s.bookRepo.FindByID(ctx, bookID)
		if err != nil {
			return dtos.SetMultipleReturnedResponse{}, errors.New("book not found: " + bookID)
		}
		if book.Status == constants.STATUS_AVAILABLE {
			return dtos.SetMultipleReturnedResponse{}, errors.New("book already available: " + bookID)
		}

		borrowing, err := s.borrowingRepo.FindCurrentlyByBookID(ctx, bookID)
		if err != nil {
			return dtos.SetMultipleReturnedResponse{}, err
		}

		borrowing.ReturnedAt = &now
		updatedBorrowing, err := s.borrowingRepo.Update(ctx, nil, borrowing)
		if err != nil {
			return dtos.SetMultipleReturnedResponse{}, err
		}

		book.Status = constants.STATUS_AVAILABLE
		book.UpdatedAt = now
		updatedBook, err := s.bookRepo.Update(ctx, nil, book)
		if err != nil {
			return dtos.SetMultipleReturnedResponse{}, err
		}

		results = append(results, dtos.SetBookReturnedResponse{
			BookID:      updatedBook.BookID,
			BorrowingID: updatedBorrowing.BorrowingID,
			ReturnedAt:  *updatedBorrowing.ReturnedAt,
		})
	}

	return dtos.SetMultipleReturnedResponse{Returned: results}, nil
}

func (s *bookService) Delete(ctx context.Context, bookID string) error {
	mu.Lock()
	defer mu.Unlock()

	book, err := s.bookRepo.FindByID(ctx, bookID)
	if err != nil {
		return err
	}

	if book.Status == constants.STATUS_BORROWED {
		return errors.New("book is currently borrowed")
	}

	return s.bookRepo.Delete(ctx, nil, bookID)
}
