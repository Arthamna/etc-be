package services

import (
	"context"
	"errors"
	"etc-backend/constants"
	"etc-backend/internal/dtos"
	"etc-backend/internal/models"
	"etc-backend/internal/repositories"
	"fmt"
	"math"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type RekrutmenService interface {
	Create(ctx context.Context, userID string, req dtos.CreateRekrutmenRequest) (*dtos.RekrutmenResponse, error)
	GetAll(ctx context.Context, page, limit int, kegiatan, role, keyword string) (*dtos.RekrutmenListResponse, error)
	GetByID(ctx context.Context, id string) (*dtos.RekrutmenResponse, error)
	GetByUserID(ctx context.Context, userID string) ([]dtos.RekrutmenResponse, error)
	Update(ctx context.Context, userID, rekrutmenID string, req dtos.CreateRekrutmenRequest) (*dtos.RekrutmenResponse, error)
	Delete(ctx context.Context, userID, rekrutmenID string) error
	Apply(ctx context.Context, userID, rekrutmenID string, req dtos.ApplyRequest) (*dtos.PendaftarResponse, error)
	UploadCV(ctx context.Context, userID, rekrutmenID string, file *multipart.FileHeader) (*dtos.UploadFileResponse, error)
	UploadPortfolio(ctx context.Context, userID, rekrutmenID string, file *multipart.FileHeader) (*dtos.UploadFileResponse, error)
	UpdateApplyStatus(ctx context.Context, rekrutmenID, pendaftarID, status string) error
	GetApplicantDetail(ctx context.Context, rekrutmenID, pendaftarID string) (*dtos.PendaftarResponse, error)
	GetTeamMembers(ctx context.Context, timID string) ([]dtos.TimMemberResponse, error)
	GiveMemberRating(ctx context.Context, reviewerID, timID, targetUserID string, rating int64, deskripsi string) error
}

type rekrutmenService struct {
    rekrutmenRepo repositories.RekrutmenRepository
    pendaftarRepo repositories.PendaftarRepository
    timRepo       repositories.TimRepository
    historyRepo   repositories.HistoryRepository
    driveService  SettingDriveService
}

func NewRekrutmenService(
    rekrutmenRepo repositories.RekrutmenRepository,
    pendaftarRepo repositories.PendaftarRepository,
    timRepo repositories.TimRepository,
    historyRepo repositories.HistoryRepository,
    driveService SettingDriveService,
) RekrutmenService {
    return &rekrutmenService{
        rekrutmenRepo: rekrutmenRepo,
        pendaftarRepo: pendaftarRepo,
        timRepo:       timRepo,
        historyRepo:   historyRepo,
        driveService:  driveService,
    }
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

func (s *rekrutmenService) Apply(ctx context.Context, userID, rekrutmenID string, req dtos.ApplyRequest) (*dtos.PendaftarResponse, error) {
    existing, _ := s.pendaftarRepo.FindByUserAndRekrutmen(ctx, userID, rekrutmenID)
    if existing != nil {
        return nil, errors.New("sudah mendaftar ke rekrutmen ini")
    }

    pendaftar := &models.Pendaftar{
        PendaftarID:     uuid.NewString(),
        RekrutmenID:     rekrutmenID,
        UserID:          userID,
        CVURL:           req.CVURL,
        PortofolioURL:   req.PortofolioURL,
        Status:          constants.STATUS_PENDING,
    }

    if req.AlasanMendaftar != "" {
        pendaftar.AlasanMendaftar = &req.AlasanMendaftar
    }

    created, err := s.pendaftarRepo.Create(ctx, pendaftar)
    if err != nil {
        return nil, err
    }

    created, err = s.pendaftarRepo.FindByID(ctx, created.PendaftarID)
    if err != nil {
        return nil, err
    }

    return &dtos.PendaftarResponse{
        PendaftarID:     created.PendaftarID,
        RekrutmenID:     created.RekrutmenID,
        UserID:          created.UserID,
        AlasanMendaftar: created.AlasanMendaftar,
        CVURL:           created.CVURL,
        PortofolioURL:   created.PortofolioURL,
        Status:          created.Status,
        NamaPendaftar:   created.User.Nama,
    }, nil
}

func (s *rekrutmenService) UploadCV(ctx context.Context, userID, rekrutmenID string, file *multipart.FileHeader) (*dtos.UploadFileResponse, error) {
    url, err := s.driveService.UploadFile(constants.DRIVE_REKRUTMEN, userID+"_cv", file)
    if err != nil {
        return nil, err
    }
    link := fmt.Sprintf("https://drive.google.com/file/d/%s/view", url)
    return &dtos.UploadFileResponse{URL: link}, nil
}

func (s *rekrutmenService) UploadPortfolio(ctx context.Context, userID, rekrutmenID string, file *multipart.FileHeader) (*dtos.UploadFileResponse, error) {
    url, err := s.driveService.UploadFile(constants.DRIVE_REKRUTMEN, userID+"_portfolio", file)
    if err != nil {
        return nil, err
    }
    link := fmt.Sprintf("https://drive.google.com/file/d/%s/view", url)
    return &dtos.UploadFileResponse{URL: link}, nil
}

func (s *rekrutmenService) UpdateApplyStatus(ctx context.Context, rekrutmenID, pendaftarID, status string) error {
    pendaftar, err := s.pendaftarRepo.FindByID(ctx, pendaftarID)
    if err != nil {
        return errors.New("pendaftar tidak ditemukan")
    }

    if pendaftar.RekrutmenID != rekrutmenID {
        return errors.New("pendaftar tidak sesuai dengan rekrutmen")
    }

    if status == constants.STATUS_APPROVED {
        tim, err := s.timRepo.FindByRekrutmenID(ctx, rekrutmenID)
        if err != nil {
            rekrutmen, err := s.rekrutmenRepo.FindByID(ctx, rekrutmenID)
            if err != nil {
                return err
            }
            tim = &models.Tim{
                TimID:       uuid.NewString(),
                TipeTim:     rekrutmen.Kegiatan,
                RekrutmenID: rekrutmenID,
                NamaKetua:   rekrutmen.User.Nama,
                CreatedAt:   time.Now(),
            }
            if err := s.timRepo.Create(ctx, tim); err != nil {
                return err
            }
        }

        count, _ := s.timRepo.CountParticipants(ctx, tim.TimID)
        participant := &models.TimParticipant{
            ID:       uuid.NewString(),
            TimID:    tim.TimID,
            UserID:   pendaftar.UserID,
            MemberKe: count + 1,
        }
        if err := s.timRepo.AddParticipant(ctx, participant); err != nil {
            return err
        }
    }

    return s.pendaftarRepo.UpdateStatus(ctx, pendaftarID, status)
}

func (s *rekrutmenService) GetApplicantDetail(ctx context.Context, rekrutmenID, pendaftarID string) (*dtos.PendaftarResponse, error) {
    pendaftar, err := s.pendaftarRepo.FindByID(ctx, pendaftarID)
    if err != nil {
        return nil, errors.New("pendaftar tidak ditemukan")
    }

    return &dtos.PendaftarResponse{
        PendaftarID:     pendaftar.PendaftarID,
        RekrutmenID:     pendaftar.RekrutmenID,
        UserID:          pendaftar.UserID,
        AlasanMendaftar: pendaftar.AlasanMendaftar,
        CVURL:           pendaftar.CVURL,
        PortofolioURL:   pendaftar.PortofolioURL,
        Status:          pendaftar.Status,
        NamaPendaftar:   pendaftar.User.Nama,
    }, nil
}

func (s *rekrutmenService) GetTeamMembers(ctx context.Context, timID string) ([]dtos.TimMemberResponse, error) {
    tim, err := s.timRepo.FindByID(ctx, timID)
    if err != nil {
        return nil, errors.New("tim tidak ditemukan")
    }

    var responses []dtos.TimMemberResponse
    for _, p := range tim.Participants {
        jurusan := ""
        if p.User.Jurusan != nil {
            jurusan = *p.User.Jurusan
        }
        responses = append(responses, dtos.TimMemberResponse{
            UserID:   p.UserID,
            Nama:     p.User.Nama,
            Jurusan:  jurusan,
            MemberKe: p.MemberKe,
        })
    }
    return responses, nil
}

func (s *rekrutmenService) GiveMemberRating(ctx context.Context, reviewerID, timID, targetUserID string, rating int64, deskripsi string) error {
    existing, _ := s.historyRepo.FindByReviewerAndUserAndTim(ctx, reviewerID, targetUserID, timID)
    if existing != nil {
        return errors.New("sudah memberikan rating ke user ini")
    }

    history := &models.History{
        ID:             uuid.NewString(),
        UserID:         targetUserID,
        ReviewerUserID: reviewerID,
        TimID:          timID,
        Rating:         rating,
        Deskripsi:      deskripsi,
        CreatedAt:      time.Now(),
    }

    return s.historyRepo.Create(ctx, history)
}
