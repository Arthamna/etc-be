package services

import (
	"arthamna/rplLibrary/internal/dtos"
	"arthamna/rplLibrary/internal/models"
	"arthamna/rplLibrary/internal/repositories"
	"context"

	"github.com/google/uuid"
)

type (
	CategoryService interface {
		Create(ctx context.Context, req dtos.CategoryCreateRequest) (dtos.CategoryResponse, error)
		GetAll(ctx context.Context) ([]dtos.CategoryResponse, error)
		GetByID(ctx context.Context, id string) (dtos.CategoryResponse, error)
		Update(ctx context.Context, id string, req dtos.CategoryUpdateRequest) (dtos.CategoryResponse, error)
		Delete(ctx context.Context, id string) error
	}

	categoryService struct {
		categoryRepo repositories.CategoryRepository
	}
)

func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) Create(ctx context.Context, req dtos.CategoryCreateRequest) (dtos.CategoryResponse, error) {
	category := &models.Category{
		CategoryID:  uuid.NewString(),
		Name:        req.Name,
		Description: req.Description,
	}

	createdCategory, err := s.categoryRepo.Create(ctx, nil, category)
	if err != nil {
		return dtos.CategoryResponse{}, err
	}

	return dtos.CategoryResponse{
		CategoryID:  createdCategory.CategoryID,
		Name:        createdCategory.Name,
		Description: createdCategory.Description,
	}, nil
}

func (s *categoryService) GetAll(ctx context.Context) ([]dtos.CategoryResponse, error) {
	categories, err := s.categoryRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var response []dtos.CategoryResponse
	for _, category := range categories {
		category := category
		response = append(response, dtos.CategoryResponse{
			CategoryID:  category.CategoryID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return response, nil
}

func (s *categoryService) GetByID(ctx context.Context, id string) (dtos.CategoryResponse, error) {
	category, err := s.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return dtos.CategoryResponse{}, err
	}

	return dtos.CategoryResponse{
		CategoryID:  category.CategoryID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (s *categoryService) Update(ctx context.Context, id string, req dtos.CategoryUpdateRequest) (dtos.CategoryResponse, error) {
	category, err := s.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return dtos.CategoryResponse{}, err
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Description != "" {
		category.Description = req.Description
	}

	updatedCategory, err := s.categoryRepo.Update(ctx, nil, category)
	if err != nil {
		return dtos.CategoryResponse{}, err
	}

	return dtos.CategoryResponse{
		CategoryID:  updatedCategory.CategoryID,
		Name:        updatedCategory.Name,
		Description: updatedCategory.Description,
	}, nil
}

func (s *categoryService) Delete(ctx context.Context, id string) error {
	if _, err := s.categoryRepo.FindByID(ctx, id); err != nil {
		return err
	}

	return s.categoryRepo.Delete(ctx, nil, id)
}
