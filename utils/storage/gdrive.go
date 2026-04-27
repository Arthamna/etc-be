package storage

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"strings"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type (
	Gdrive interface {
		UploadFile(filename string, content io.Reader, parentId string) (string, error)
		MakeFolder(foldername string, parentId string) (string, error)
		SetPermissionToDownload(fileId string) error
		SetPermission(fileId string, typePermission string, rolePermission string) error
	}

	gdrive struct {
		gdriveService *drive.Service
	}
)

func NewGdrive() Gdrive {
	driveCredentialsFile := map[string]interface{}{
		"type":                        os.Getenv("DRIVE_TYPE"),
		"project_id":                  os.Getenv("DRIVE_PROJECT_ID"),
		"private_key_id":              os.Getenv("DRIVE_PRIVATE_KEY_ID"),
		"private_key":                 strings.Replace(os.Getenv("DRIVE_PRIVATE_KEY"), "/\\n/gm", "\n", -1),
		"client_email":                os.Getenv("DRIVE_CLIENT_EMAIL"),
		"client_id":                   os.Getenv("DRIVE_CLIENT_ID"),
		"auth_uri":                    os.Getenv("DRIVE_AUTH_URI"),
		"token_uri":                   os.Getenv("DRIVE_TOKEN_URI"),
		"auth_provider_x509_cert_url": os.Getenv("DRIVE_AUTH_PROVIDER_x509_CERT_URL"),
		"client_x509_cert_url":        os.Getenv("DRIVE_CLIENT_x509_CERT_URL"),
		"universe_domain":             os.Getenv("DRIVE_UNIVERSE_DOMAIN"),
	}

	credentials, err := json.Marshal(driveCredentialsFile)
	if err != nil {
		panic("Failed connect to google drive api: " + err.Error())
	}

	ctx := context.Background()
	opt := option.WithCredentialsJSON(credentials)

	gdriveService, err := drive.NewService(ctx, opt)
	if err != nil {
		panic("Failed connect to google drive api: " + err.Error())
	}

	return &gdrive{
		gdriveService: gdriveService,
	}
}

func (gs *gdrive) UploadFile(filename string, content io.Reader, parentId string) (string, error) {
	f := &drive.File{
		Name:    filename,
		Parents: []string{parentId},
	}

	file, err := gs.gdriveService.Files.Create(f).Media(content).Do()
	if err != nil {
		return "", err
	}

	return file.Id, nil
}

func (gs *gdrive) MakeFolder(foldername string, parentId string) (string, error) {
	folder := &drive.File{
		Name:     foldername,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{parentId},
	}

	file, err := gs.gdriveService.Files.Create(folder).Do()
	if err != nil {
		return "", err
	}

	return file.Id, nil
}

func (gs *gdrive) SetPermissionToDownload(fileId string) error {
	return gs.SetPermission(fileId, "anyone", "reader")
}

func (gs *gdrive) SetPermission(fileId string, typePermission string, rolePermission string) error {
	permission := &drive.Permission{
		Type: typePermission,
		Role: rolePermission,
	}

	_, err := gs.gdriveService.Permissions.Create(fileId, permission).Do()
	if err != nil {
		return err
	}

	return nil
}
