package services

import (
	"arthamna/rplLibrary/constants"
	"arthamna/rplLibrary/internal/dtos"
	"arthamna/rplLibrary/internal/models"
	"arthamna/rplLibrary/internal/repositories"
	"arthamna/rplLibrary/services/auth"
	"context"
	"encoding/base64"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		Register(ctx context.Context, req dtos.UserRegisterRequest) (dtos.UserRegisterResponse, error)
		RegisterAdmin(ctx context.Context, req dtos.AdminRegisterRequest) (dtos.UserRegisterResponse, error)
		Login(ctx context.Context, req dtos.UserLoginRequest) (dtos.UserLoginResponse, error)
		UploadProfilePicture(ctx context.Context, req dtos.UploadProfilePictureRequest, userId string) (dtos.UpdateProfilePictureResponse, error)
	}

	userService struct {
		userRepo   repositories.UserRepository
		jwtService auth.JWTService
	}
)

func NewUserService(userRepo repositories.UserRepository, jwtService auth.JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

var mu sync.Mutex

func readImageBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, errors.New("gagal membuka file gambar")
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("gagal membaca file gambar")
	}

	return bytes, nil
}

func (s *userService) Register(ctx context.Context, req dtos.UserRegisterRequest) (dtos.UserRegisterResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	existingUser, _ := s.userRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return dtos.UserRegisterResponse{}, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dtos.UserRegisterResponse{}, err
	}

	now := time.Now()
	user := &models.User{
		UserID:           uuid.NewString(),
		Username:         req.Username,
		Email:            req.Email,
		PasswordHash:     string(hashedPassword),
		Role:             constants.ROLE_USER,
		RegistrationDate: now,
		CreatedAt:        now,
		UpdatedAt:        now,
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

	user, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		return dtos.UpdateProfilePictureResponse{}, err
	}

	imageBytes, err := readImageBytes(req.ProfilePicture)
	if err != nil {
		return dtos.UpdateProfilePictureResponse{}, err
	}

	user.ProfilePicture = imageBytes
	user.UpdatedAt = time.Now()

	updatedUser, err := s.userRepo.Update(ctx, nil, user)
	if err != nil {
		return dtos.UpdateProfilePictureResponse{}, err
	}

	return dtos.UpdateProfilePictureResponse{
		ProfilePicture: base64.StdEncoding.EncodeToString(updatedUser.ProfilePicture),
	}, nil
}

func (s *userService) Login(ctx context.Context, req dtos.UserLoginRequest) (dtos.UserLoginResponse, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return dtos.UserLoginResponse{}, errors.New("invalid email")
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

func (s *userService) RegisterAdmin(ctx context.Context, req dtos.AdminRegisterRequest) (dtos.UserRegisterResponse, error) {
	expectedKey := os.Getenv("ADMIN_SECRET_KEY")
	if expectedKey == "" {
		return dtos.UserRegisterResponse{}, errors.New("admin secret key not configured")
	}

	if req.SecretKey != expectedKey {
		return dtos.UserRegisterResponse{}, errors.New("invalid admin secret key")
	}

	existingUser, _ := s.userRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return dtos.UserRegisterResponse{}, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dtos.UserRegisterResponse{}, err
	}

	now := time.Now()
	user := &models.User{
		UserID:           uuid.New().String(),
		Username:         req.Username,
		Email:            req.Email,
		PasswordHash:     string(hashedPassword),
		Role:             constants.ROLE_ADMIN,
		RegistrationDate: now,
		CreatedAt:        now,
		UpdatedAt:        now,
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
