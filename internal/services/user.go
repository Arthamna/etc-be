package services

import (
	"context"
	"errors"
	"etc-backend/constants"
	"etc-backend/internal/dtos"
	"etc-backend/internal/models"
	"etc-backend/internal/repositories"
	// common "etc-backend/utils/response"
	"etc-backend/utils/storage"
	"fmt"

	// "os"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		Register(ctx context.Context, req dtos.UserRegisterRequest) (dtos.UserRegisterResponse, error)
		Login(ctx context.Context, req dtos.UserLoginRequest) (dtos.UserLoginResponse, error)
		UpdateUser(ctx context.Context, userID string, req dtos.UpdateUserRequest) (*dtos.UserResponse, error)
		GetBookmarks(ctx context.Context, userID string) ([]dtos.BookmarkResponse, error)
		GetByID(ctx context.Context, userId string) (dtos.UserGetMe, error)
		UploadProfilePicture(ctx context.Context, req dtos.UploadProfilePictureRequest, userId string) (dtos.UpdateProfilePictureResponse, error)
	}

	userService struct {
		userRepo     repositories.UserRepository
		jwtService   JWTService
		driveService SettingDriveService
		gdrive       storage.Gdrive
	}
)

func NewUserService(userRepo repositories.UserRepository, jwtService JWTService, driveService SettingDriveService, gdrive storage.Gdrive) UserService {
	return &userService{
		userRepo:     userRepo,
		jwtService:   jwtService,
		driveService: driveService,
		gdrive:       gdrive,
	}
}

var mu sync.Mutex

func (s *userService) Register(ctx context.Context, req dtos.UserRegisterRequest) (dtos.UserRegisterResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	// validate based on role
	if req.Role == constants.USER_NRP {
		if *req.Jurusan == "" {
			return dtos.UserRegisterResponse{}, errors.New("mahasiswa harus mengisi jurusan")
		}
		existing, _ := s.userRepo.FindByNoPengenal(ctx, req.NoPengenal)
		if existing != nil {
			return dtos.UserRegisterResponse{}, errors.New("nrp already registered")
		}
	} else if req.Role == constants.USER_NIDN {
		existing, _ := s.userRepo.FindByNoPengenal(ctx, req.NoPengenal)
		if existing != nil {
			return dtos.UserRegisterResponse{}, errors.New("nidn already registered")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dtos.UserRegisterResponse{}, err
	}

	now := time.Now()
	user := &models.User{
		UserID:       uuid.NewString(),
		Nama:         req.Nama,
		Role:         req.Role,
		NoTelp:       req.NoTelp,
		PasswordHash: string(hashedPassword),
		CreatedAt:    now,
		UpdatedAt:    now,
		Spesialisasi: req.Spesialisasi,
	}

	if req.Role == "mahasiswa" {
		user.NoPengenal = req.NoPengenal
		user.Jurusan = req.Jurusan
	} else if req.Role == "dosen" {
		user.NoPengenal = req.NoPengenal
	}

	createdUser, err := s.userRepo.Create(ctx, nil, user)
	if err != nil {
		return dtos.UserRegisterResponse{}, err
	}

	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return dtos.UserRegisterResponse{}, err
	}

	return dtos.UserRegisterResponse{
		User:  dtos.UserResponse{
			UserID:   createdUser.UserID,
			Nama:     createdUser.Nama,
			Role:     createdUser.Role,
			NoTelp:   createdUser.NoTelp,
			NoPengenal:      createdUser.NoPengenal,
			Jurusan:  createdUser.Jurusan,
			Spesialisasi: createdUser.Spesialisasi,
		},
		Token: token,
	}, nil
}

func (s *userService) UploadProfilePicture(ctx context.Context, req dtos.UploadProfilePictureRequest, userId string) (dtos.UpdateProfilePictureResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	if req.ProfilePicture == nil {
		return dtos.UpdateProfilePictureResponse{}, errors.New("profile picture is required")
	}

	// filename := userId
	url, err := s.driveService.UploadFile(constants.DRIVE_USER, userId, req.ProfilePicture)

	user, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		return dtos.UpdateProfilePictureResponse{}, err
	}

	link := fmt.Sprintf("https://drive.google.com/file/d/%s/view", url)
	user.ProfilePicture = &link
	user.UpdatedAt = time.Now()

	updatedUser, err := s.userRepo.Update(ctx, nil, user)
	if err != nil {
		return dtos.UpdateProfilePictureResponse{}, err
	}

	return dtos.UpdateProfilePictureResponse{
		ProfilePicture: updatedUser.ProfilePicture,
	}, nil
}

func (s *userService) GetByID(ctx context.Context, userId string) (dtos.UserGetMe, error) {
	user, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		return dtos.UserGetMe{}, errors.New("user not found")
	}


	return dtos.UserGetMe{
		PersonalInfo: dtos.UserResponse{
			UserID:         user.UserID,
			Nama:           user.Nama,
			Jurusan:        user.Jurusan,
			NoPengenal: user.NoPengenal,
			NoTelp:         user.NoTelp,
			Role:           user.Role,
			ProfilePicture: user.ProfilePicture,
		},
	}, nil
}

func (s *userService) Login(ctx context.Context, req dtos.UserLoginRequest) (dtos.UserLoginResponse, error) {
	// var user *models.User
	var err error

	user, err := s.userRepo.FindByNoPengenal(ctx, req.NoPengenal)
	if err != nil {
		return dtos.UserLoginResponse{}, errors.New("data tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return dtos.UserLoginResponse{}, errors.New("credentials salah")
	}

	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return dtos.UserLoginResponse{}, err
	}

	return dtos.UserLoginResponse{
		Token: token,
		Role:  user.Role,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, userID string, req dtos.UpdateUserRequest) (*dtos.UserResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.Nama != "" {
		user.Nama = req.Nama
	}
	if req.Jurusan != "" {
		user.Jurusan = &req.Jurusan
	}
	if req.NoTelp != "" {
		user.NoTelp = req.NoTelp
	}
	if req.Spesialisasi != "" {
		user.Spesialisasi = []string{req.Spesialisasi}
	}

	user.UpdatedAt = time.Now()

	updated, err := s.userRepo.Update(ctx, nil, user)
	if err != nil {
		return nil, err
	}

	return &dtos.UserResponse{
		UserID:         updated.UserID,
		Nama:           updated.Nama,
		Jurusan:        updated.Jurusan,
		NoPengenal:           updated.NoPengenal,
		NoTelp:         updated.NoTelp,
		Role:           updated.Role,
		ProfilePicture: updated.ProfilePicture,
	}, nil
}

func (s *userService) GetBookmarks(ctx context.Context, userID string) ([]dtos.BookmarkResponse, error) {
	bookmarks, err := s.userRepo.FindBookmarksByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responses []dtos.BookmarkResponse
	for _, b := range bookmarks {
		responses = append(responses, dtos.BookmarkResponse{
			ID:          b.ID,
			RekrutmenID: b.RekrutmenID,
			Rekrutmen:   dtos.ToRekrutmenResponse(&b.Rekrutmen),
		})
	}
	return responses, nil
}
