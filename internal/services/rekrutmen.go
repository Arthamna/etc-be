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
	GetAppliedByID(ctx context.Context, id string) (*dtos.ApplierRekrutmenResponse, error)
	GetByUserID(ctx context.Context, userID string) ([]dtos.RekrutmenResponse, error)
	Update(ctx context.Context, userID, rekrutmenID string, req dtos.CreateRekrutmenRequest) (*dtos.RekrutmenResponse, error)
	Delete(ctx context.Context, userID, rekrutmenID string) error
	Apply(ctx context.Context, userID, rekrutmenID string, req dtos.ApplyRequest) (*dtos.PendaftarResponse, error)
	UploadCV(ctx context.Context, userID, rekrutmenID string, file *multipart.FileHeader) (*dtos.UploadFileResponse, error)
	UploadPortfolio(ctx context.Context, userID, rekrutmenID string, file *multipart.FileHeader) (*dtos.UploadFileResponse, error)
    AcceptPendaftar(ctx context.Context, userID, rekrutmenID, pendaftarID string) error
    RejectPendaftar(ctx context.Context, userID, rekrutmenID, pendaftarID string) error
	RefreshApplyStatus(ctx context.Context, rekrutmenID, pendaftarID, status string) error
	GetAppliedRekrutmen(ctx context.Context, userID string) ([]dtos.MyApplicationResponse, error)
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
		UserID:         userID,
		Kegiatan:       req.Kegiatan,
        Kriteria:       req.Kriteria,
		TanggalMulai:   req.TanggalMulai,
		TanggalSelesai: req.TanggalSelesai,
		Fee:            &req.Fee,
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

	res := dtos.RekrutmenResponse{
        RekrutmenID:    created.RekrutmenID,
        UserID:         created.UserID,
        Kegiatan:       created.Kegiatan,
        Kriteria:       created.Kriteria,
        TanggalMulai:   created.TanggalMulai,
        TanggalSelesai: created.TanggalSelesai,
        Fee:            *created.Fee,
        Role:           created.Role,
        ContactPerson:  created.ContactPerson,
    }
	return &res, nil
}

func (s *rekrutmenService) GetByID(ctx context.Context, id string) (*dtos.RekrutmenResponse, error) {
    data, err := s.rekrutmenRepo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }

    timID := ""
    if len(data.Tims) > 0 {
        timID = data.Tims[0].TimID
    }

    res := dtos.RekrutmenResponse{
        RekrutmenID:    data.RekrutmenID,
        UserID:         data.UserID,
        Kegiatan:       data.Kegiatan,
        Kriteria:       data.Kriteria,
        TanggalMulai:   data.TanggalMulai,
        TanggalSelesai: data.TanggalSelesai,
        Fee:            *data.Fee,
        Role:           data.Role,
        ContactPerson:  data.ContactPerson,
        TimID:          timID,
    }
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
        timID := ""
        if len(r.Tims) > 0 {
            timID = r.Tims[0].TimID
        }
        responses = append(responses, dtos.RekrutmenResponse{
            RekrutmenID:    r.RekrutmenID,
            UserID:         r.UserID,
            Kegiatan:       r.Kegiatan,
            Kriteria:       r.Kriteria,
            TanggalMulai:   r.TanggalMulai,
            TanggalSelesai: r.TanggalSelesai,
            Fee:            *r.Fee,
            Role:           r.Role,
            ContactPerson:  r.ContactPerson,
            TimID:          timID,
        })
	}

	return &dtos.RekrutmenListResponse{
		Data:       responses,
		Page:       page,
		Limit:      limit,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
	}, nil
}

func (s *rekrutmenService) GetAppliedByID(ctx context.Context, id string) (*dtos.ApplierRekrutmenResponse, error) {
	rekrutmen, err := s.rekrutmenRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("rekrutmen not found")
	}

    pendaftar, err := s.pendaftarRepo.FindByRekrutmen(ctx, id)
    if err != nil {
        return nil, err
    }

    timID := ""
    if len(rekrutmen.Tims) > 0 {
        timID = rekrutmen.Tims[0].TimID
    }

    res := dtos.ApplierRekrutmenResponse{
        RekrutmenID:    rekrutmen.RekrutmenID,
        UserID:         rekrutmen.UserID,
        Kegiatan:       rekrutmen.Kegiatan,
        Kriteria:       rekrutmen.Kriteria,
        TanggalMulai:   rekrutmen.TanggalMulai,
        TanggalSelesai: rekrutmen.TanggalSelesai,
        Fee:            *rekrutmen.Fee,
        Role:           rekrutmen.Role,
        ContactPerson:  rekrutmen.ContactPerson,
        TimID:          timID,
    }
    for _, p := range *pendaftar {
        res.Pendaftar = append(res.Pendaftar, dtos.PendaftarResponse{
            PendaftarID:     p.PendaftarID,
            RekrutmenID:     p.RekrutmenID,
            UserID:          p.UserID,
            AlasanMendaftar: p.AlasanMendaftar,
            CVURL:           p.CVURL,
            PortofolioURL:   p.PortofolioURL,
            Status:          p.Status,
            NamaPendaftar:   p.User.Nama,
        })
    }
    return &res, nil
}

func (s *rekrutmenService) GetByUserID(ctx context.Context, userID string) ([]dtos.RekrutmenResponse, error) {
	data, err := s.rekrutmenRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responses []dtos.RekrutmenResponse
	for _, r := range data {
        timID := ""
        if len(r.Tims) > 0 {
            timID = r.Tims[0].TimID
        }
		responses = append(responses, dtos.RekrutmenResponse{
            RekrutmenID:    r.RekrutmenID,
            UserID:         r.UserID,
            Kegiatan:       r.Kegiatan,
            Kriteria:       r.Kriteria,
            TanggalMulai:   r.TanggalMulai,
            TanggalSelesai: r.TanggalSelesai,
            Fee:            *r.Fee,
            Role:           r.Role,
            ContactPerson:  r.ContactPerson,
            TimID:          timID,
        })
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

    if req.Kegiatan != "" {
        rekrutmen.Kegiatan = req.Kegiatan
    }
    if !req.TanggalMulai.IsZero() {
        rekrutmen.TanggalMulai = req.TanggalMulai
    }
    if !req.TanggalSelesai.IsZero() {
        rekrutmen.TanggalSelesai = req.TanggalSelesai
    }
    if req.Fee != 0 {
        rekrutmen.Fee = &req.Fee
    }
    if req.Role != "" {
        rekrutmen.Role = req.Role
    }
    if req.ContactPerson != "" {
        rekrutmen.ContactPerson = req.ContactPerson
    }

	updated, err := s.rekrutmenRepo.Update(ctx, rekrutmen)
	if err != nil {
		return nil, err
	}

	updated, err = s.rekrutmenRepo.FindByID(ctx, updated.RekrutmenID)
	if err != nil {
		return nil, err
	}

	res := dtos.RekrutmenResponse{
        RekrutmenID:    rekrutmen.RekrutmenID,
        UserID:         rekrutmen.UserID,
        Kegiatan:       rekrutmen.Kegiatan,
        Kriteria:       rekrutmen.Kriteria,
        TanggalMulai:   rekrutmen.TanggalMulai,
        TanggalSelesai: rekrutmen.TanggalSelesai,
        Fee:            *rekrutmen.Fee,
        Role:           rekrutmen.Role,
        ContactPerson:  rekrutmen.ContactPerson,
    }
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

func (s *rekrutmenService) RefreshApplyStatus(ctx context.Context, rekrutmenID, pendaftarID, status string) error {
	pendaftar, err := s.pendaftarRepo.FindByID(ctx, pendaftarID)
	if err != nil {
		return errors.New("pendaftar tidak ditemukan")
	}

	if pendaftar.RekrutmenID != rekrutmenID {
		return errors.New("pendaftar tidak sesuai dengan rekrutmen")
	}

	return s.handleStatusUpdate(ctx, rekrutmenID, pendaftarID, pendaftar.UserID, status)
}

func (s *rekrutmenService) AcceptPendaftar(ctx context.Context, userID, rekrutmenID, pendaftarID string) error {
	ok, err := s.rekrutmenRepo.IsOwnedByUser(ctx, rekrutmenID, userID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("rekrutmen bukan milik user ini")
	}

	pendaftar, err := s.pendaftarRepo.FindByID(ctx, pendaftarID)
	if err != nil {
		return errors.New("pendaftar tidak ditemukan")
	}

	return s.handleStatusUpdate(ctx, rekrutmenID, pendaftarID, pendaftar.UserID, constants.STATUS_APPROVED)
}

func (s *rekrutmenService) RejectPendaftar(ctx context.Context, userID, rekrutmenID, pendaftarID string) error {
	ok, err := s.rekrutmenRepo.IsOwnedByUser(ctx, rekrutmenID, userID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("rekrutmen bukan milik user ini")
	}

	pendaftar, err := s.pendaftarRepo.FindByID(ctx, pendaftarID)
	if err != nil {
		return errors.New("pendaftar tidak ditemukan")
	}

	return s.handleStatusUpdate(ctx, rekrutmenID, pendaftarID, pendaftar.UserID, constants.STATUS_REJECTED)
}

func (s *rekrutmenService) handleStatusUpdate(ctx context.Context, rekrutmenID, pendaftarID, applicantUserID, status string) error {
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

		// Check if already a participant to avoid duplicates
		isMember, err := s.timRepo.IsParticipant(ctx, tim.TimID, applicantUserID)
		if err != nil {
			return err
		}

		if !isMember {
			count, _ := s.timRepo.CountParticipants(ctx, tim.TimID)
			participant := &models.TimParticipant{
				ID:       uuid.NewString(),
				TimID:    tim.TimID,
				UserID:   applicantUserID,
				MemberKe: count + 1,
			}
			if err := s.timRepo.AddParticipant(ctx, participant); err != nil {
				return err
			}
		}
	} else if status == constants.STATUS_REJECTED {
		// If rejected, remove from team if they were previously added
		tim, err := s.timRepo.FindByRekrutmenID(ctx, rekrutmenID)
		if err == nil && tim != nil {
			if err := s.timRepo.RemoveParticipant(ctx, tim.TimID, applicantUserID); err != nil {
				return err
			}
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

func (s *rekrutmenService) GetAppliedRekrutmen(ctx context.Context, userID string) ([]dtos.MyApplicationResponse, error) {
    data, err := s.pendaftarRepo.FindByUserID(ctx, userID)
    if err != nil {
        return nil, err
    }

    var responses []dtos.MyApplicationResponse
    for _, p := range data {
        timID := ""
        if p.Status == constants.STATUS_APPROVED && len(p.Rekrutmen.Tims) > 0 {
            timID = p.Rekrutmen.Tims[0].TimID
        }

        fee := 0.0
        if p.Rekrutmen.Fee != nil {
            fee = *p.Rekrutmen.Fee
        }

        responses = append(responses, dtos.MyApplicationResponse{
            PendaftarID: p.PendaftarID,
            Status:      p.Status,
            Rekrutmen: dtos.RekrutmenResponse{
                RekrutmenID:    p.Rekrutmen.RekrutmenID,
                UserID:         p.Rekrutmen.UserID,
                Kegiatan:       p.Rekrutmen.Kegiatan,
                Kriteria:       p.Rekrutmen.Kriteria,
                TanggalMulai:   p.Rekrutmen.TanggalMulai,
                TanggalSelesai: p.Rekrutmen.TanggalSelesai,
                Fee:            fee,
                Role:           p.Rekrutmen.Role,
                ContactPerson:  p.Rekrutmen.ContactPerson,
                TimID:          timID,
            },
        })
    }
    return responses, nil
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
			UserID:       p.UserID,
			Nama:         p.User.Nama,
			Jurusan:      jurusan,
			MemberKe:     p.MemberKe,
			NoPengenal:   p.User.NoPengenal,
			NoTelp:       p.User.NoTelp,
			Spesialisasi: p.User.Spesialisasi,
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
