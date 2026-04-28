package services

import (
	"context"
	"errors"
	"etc-backend/constants"
	"etc-backend/internal/dtos"
	"etc-backend/internal/models"
	"etc-backend/internal/repositories"
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
		// RegisterAdmin(ctx context.Context, req dtos.AdminRegisterRequest) (dtos.UserRegisterResponse, error)
		Login(ctx context.Context, req dtos.UserLoginRequest) (dtos.UserLoginResponse, error)
		UpdateUser(ctx context.Context, userID string, req dtos.UpdateUserRequest) (*dtos.UserResponse, error)
		GetBookmarks(ctx context.Context, userID string) ([]dtos.BookmarkResponse, error)
		GetByID(ctx context.Context, userId string) (dtos.UserGetMe, error)
		UploadProfilePicture(ctx context.Context, req dtos.UploadProfilePictureRequest, userId string) (dtos.UpdateProfilePictureResponse, error)
	}

	userService struct {
		userRepo   repositories.UserRepository
		jwtService JWTService
		driveService SettingDriveService
		gdrive		storage.Gdrive
	}
)

func NewUserService(userRepo repositories.UserRepository, jwtService JWTService, driveService SettingDriveService, gdrive storage.Gdrive) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
		driveService: driveService,
		gdrive: gdrive,
	}
}

var mu sync.Mutex

func (s *userService) Register(ctx context.Context, req dtos.UserRegisterRequest) (dtos.UserRegisterResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	existingUser, _ := s.userRepo.FindByNRP(ctx, req.NRP)
	if existingUser != nil {
		return dtos.UserRegisterResponse{}, errors.New("nrp already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dtos.UserRegisterResponse{}, err
	}

	now := time.Now()
	user := &models.User{
		UserID:        uuid.NewString(),
		Nama:          req.Nama,
		Jurusan:       req.Jurusan,
		NRP:           &req.NRP,
		NoTelp:        req.ContactPerson,
		PasswordHash:  string(hashedPassword),
		Role:          constants.ROLE_USER,
		CreatedAt:     now,
		UpdatedAt:     now,
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
		User:  *dtos.ToUserResponse(createdUser),
		Token: token,
	}, nil
}

func (s *userService) UploadProfilePicture(ctx context.Context, req dtos.UploadProfilePictureRequest, userId string) (dtos.UpdateProfilePictureResponse, error) {
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

	nrp := ""
	if user.NRP != nil {
		nrp = *user.NRP
	}

	return dtos.UserGetMe{
		PersonalInfo: dtos.UserResponse{
			UserID:        user.UserID,
			Nama:          user.Nama,
			Jurusan:       user.Jurusan,
			NRP:           nrp,
			ContactPerson: user.NoTelp,
			Role:          user.Role,
			ProfilePicture: user.ProfilePicture,
		},
	}, nil
}

func (s *userService) Login(ctx context.Context, req dtos.UserLoginRequest) (dtos.UserLoginResponse, error) {
	user, err := s.userRepo.FindByNRP(ctx, req.NRP)
	if err != nil {
		return dtos.UserLoginResponse{}, errors.New("invalid nrp")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return dtos.UserLoginResponse{}, errors.New("invalid password")
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

// func (s *userService) RegisterAdmin(ctx context.Context, req dtos.AdminRegisterRequest) (dtos.UserRegisterResponse, error) {
// 	expectedKey := os.Getenv("ADMIN_SECRET_KEY")
// 	if expectedKey == "" {
// 		return dtos.UserRegisterResponse{}, errors.New("admin secret key not configured")
// 	}

// 	if req.SecretKey != expectedKey {
// 		return dtos.UserRegisterResponse{}, errors.New("invalid admin secret key")
// 	}

// 	existingUser, _ := s.userRepo.FindByNRP(ctx, req.NRP)
// 	if existingUser != nil {
// 		return dtos.UserRegisterResponse{}, errors.New("nrp already registered")
// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return dtos.UserRegisterResponse{}, err
// 	}

// 	now := time.Now()
// 	user := &models.User{
// 		UserID:        uuid.New().String(),
// 		Nama:          req.Nama,
// 		Jurusan:       req.Jurusan,
// 		NRP:           req.NRP,
// 		ContactPerson: req.ContactPerson,
// 		PasswordHash:  string(hashedPassword),
// 		Role:          constants.ROLE_ADMIN,
// 		CreatedAt:     now,
// 		UpdatedAt:     now,
// 	}

// 	createdUser, err := s.userRepo.Create(ctx, nil, user)
// 	if err != nil {
// 		return dtos.UserRegisterResponse{}, err
// 	}

// 	token, err := s.jwtService.GenerateToken(user)
// 	if err != nil {
// 		return dtos.UserRegisterResponse{}, err
// 	}

// 	return dtos.UserRegisterResponse{
// 		User:  *dtos.ToUserResponse(createdUser),
// 		Token: token,
// 	}, nil
// }

func (s *userService) UpdateUser(ctx context.Context, userID string, req dtos.UpdateUserRequest) (*dtos.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.Nama != "" {
		user.Nama = req.Nama
	}
	if req.Jurusan != "" {
		user.Jurusan = req.Jurusan
	}
	if req.ContactPerson != "" {
		user.NoTelp = req.ContactPerson
	}
	if req.Role != "" {
		if req.Role != "mahasiswa" && req.Role != "dosen" {
			return nil, errors.New("role harus mahasiswa atau dosen")
		}
		user.Role = req.Role
	}

	user.UpdatedAt = time.Now()

	updated, err := s.userRepo.Update(ctx, nil, user)
	if err != nil {
		return nil, err
	}

	return dtos.ToUserResponse(updated), nil
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
