package services

import (
	"context"
	"errors"
	"etc-backend/internal/dtos"
	"etc-backend/internal/models"
	"etc-backend/internal/repositories"
	"math"

	"github.com/google/uuid"
)

type RekrutmenService interface {
	Create(ctx context.Context, userID string, req dtos.CreateRekrutmenRequest) (*dtos.RekrutmenResponse, error)
	GetAll(ctx context.Context, page, limit int, kegiatan, role, keyword string) (*dtos.RekrutmenListResponse, error)
	GetByID(ctx context.Context, id string) (*dtos.RekrutmenResponse, error)
	GetByUserID(ctx context.Context, userID string) ([]dtos.RekrutmenResponse, error)
	Update(ctx context.Context, userID, rekrutmenID string, req dtos.CreateRekrutmenRequest) (*dtos.RekrutmenResponse, error)
	Delete(ctx context.Context, userID, rekrutmenID string) error
}

type rekrutmenService struct {
	rekrutmenRepo repositories.RekrutmenRepository
}

func NewRekrutmenService(rekrutmenRepo repositories.RekrutmenRepository) RekrutmenService {
	return &rekrutmenService{rekrutmenRepo: rekrutmenRepo}
}

func (s *rekrutmenService) Create(ctx context.Context, userID string, req dtos.CreateRekrutmenRequest) (*dtos.RekrutmenResponse, error) {
	rekrutmen := &models.Rekrutmen{
		RekrutmenID:    uuid.NewString(),
		KegiatanID:     uuid.NewString(),
		UserID:         userID,
		Kegiatan:       req.Kegiatan,
		TanggalMulai:   req.TanggalMulai,
		TanggalSelesai: req.TanggalSelesai,
		Fee:            req.Fee,
		Role:           req.Role,
		ContactPerson:  req.ContactPerson,
	}

	created, err := s.rekrutmenRepo.Create(ctx, rekrutmen)
	if err != nil {
		return nil, err
	}

	created, err = s.rekrutmenRepo.FindByID(ctx, created.RekrutmenID)
	if err != nil {
		return nil, err
	}

	res := dtos.ToRekrutmenResponse(created)
	return &res, nil
}

func (s *rekrutmenService) GetAll(ctx context.Context, page, limit int, kegiatan, role, keyword string) (*dtos.RekrutmenListResponse, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	data, total, err := s.rekrutmenRepo.FindAll(ctx, page, limit, kegiatan, role, keyword)
	if err != nil {
		return nil, err
	}

	var responses []dtos.RekrutmenResponse
	for _, r := range data {
		responses = append(responses, dtos.ToRekrutmenResponse(&r))
	}

	return &dtos.RekrutmenListResponse{
		Data:       responses,
		Page:       page,
		Limit:      limit,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
	}, nil
}

func (s *rekrutmenService) GetByID(ctx context.Context, id string) (*dtos.RekrutmenResponse, error) {
	rekrutmen, err := s.rekrutmenRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("rekrutmen not found")
	}
	res := dtos.ToRekrutmenResponse(rekrutmen)
	return &res, nil
}

func (s *rekrutmenService) GetByUserID(ctx context.Context, userID string) ([]dtos.RekrutmenResponse, error) {
	data, err := s.rekrutmenRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responses []dtos.RekrutmenResponse
	for _, r := range data {
		responses = append(responses, dtos.ToRekrutmenResponse(&r))
	}
	return responses, nil
}

func (s *rekrutmenService) Update(ctx context.Context, userID, rekrutmenID string, req dtos.CreateRekrutmenRequest) (*dtos.RekrutmenResponse, error) {
	rekrutmen, err := s.rekrutmenRepo.FindByID(ctx, rekrutmenID)
	if err != nil {
		return nil, errors.New("rekrutmen not found")
	}

	if rekrutmen.UserID != userID {
		return nil, errors.New("tidak memiliki akses")
	}

	rekrutmen.Kegiatan = req.Kegiatan
	rekrutmen.TanggalMulai = req.TanggalMulai
	rekrutmen.TanggalSelesai = req.TanggalSelesai
	rekrutmen.Fee = req.Fee
	rekrutmen.Role = req.Role
	rekrutmen.ContactPerson = req.ContactPerson

	updated, err := s.rekrutmenRepo.Update(ctx, rekrutmen)
	if err != nil {
		return nil, err
	}

	updated, err = s.rekrutmenRepo.FindByID(ctx, updated.RekrutmenID)
	if err != nil {
		return nil, err
	}

	res := dtos.ToRekrutmenResponse(updated)
	return &res, nil
}

func (s *rekrutmenService) Delete(ctx context.Context, userID, rekrutmenID string) error {
	rekrutmen, err := s.rekrutmenRepo.FindByID(ctx, rekrutmenID)
	if err != nil {
		return errors.New("rekrutmen not found")
	}

	if rekrutmen.UserID != userID {
		return errors.New("tidak memiliki akses")
	}

	return s.rekrutmenRepo.Delete(ctx, rekrutmenID)
}
