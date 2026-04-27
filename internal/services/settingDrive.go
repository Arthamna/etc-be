package services

import (
	"etc-backend/constants"
	"etc-backend/internal/repositories"
	"etc-backend/utils/storage"
	"fmt"
	"mime/multipart"
	"os"
)

type (
	SettingDriveService interface {
		UploadFile(folderName, filename string, content *multipart.FileHeader) (string, error)
	}

	settingDriveService struct {
		settingDriveRepo repositories.SettingDriveRepository
		gdrive           storage.Gdrive
	}
)

func NewSettingDriveService(settingDriveRepo repositories.SettingDriveRepository, gdrive storage.Gdrive) SettingDriveService {
	return &settingDriveService{
		settingDriveRepo: settingDriveRepo,
		gdrive:           gdrive,
	}
}

func (s *settingDriveService) UploadFile(folderName, filename string, content *multipart.FileHeader) (string, error) {
	var targetId string
	switch folderName {
	case constants.DRIVE_REKRUTMEN:
		targetId = os.Getenv("REKRUTMEN_ID")
	case constants.DRIVE_USER:
		targetId = os.Getenv("USER_ID")
	default:
		return "", fmt.Errorf("unknown folder name: %s", folderName)
	}

	drive, err := s.settingDriveRepo.Find(targetId)
	if err != nil {
		return "", err
	}

	file, err := content.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	return s.gdrive.UploadFile(filename, file, drive.DayUrl)
}