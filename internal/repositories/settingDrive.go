package repositories

import (
	"errors"
	"etc-backend/internal/models"
	"etc-backend/utils/pagination"
	"etc-backend/utils/storage"
	"time"

	"gorm.io/gorm"
)

type (
	SettingDriveRepository interface {
		Find(targetId string) (models.SettingDrive, error)
	}

	settingDriveRepository struct {
		gdrive storage.Gdrive
		db     *gorm.DB
	}
)

func NewSettingDriveRepository(gdrive storage.Gdrive, db *gorm.DB) SettingDriveRepository {
	return &settingDriveRepository{
		gdrive: gdrive,
		db:     db,
	}
}

func (r *settingDriveRepository) Find(targetId string) (models.SettingDrive, error) {
	var settingDrive, mdrive models.SettingDrive
	cur := time.Now()
	month := cur.Month().String()
	day := cur.Day()

	// find with current month, day, and competition id to match targetid
	query := r.db.Where("month = ? AND day = ? AND type_id = ?", month, day, targetId).Take(&settingDrive)

	// if exist return
	if query.Error == nil {
		return settingDrive, nil
	}

	// if error other than record not found return it
	if !errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return models.SettingDrive{}, query.Error
	}

	// find month and target_id to get drive id
	mquery := r.db.Where("month = ? AND type_id = ?", month, targetId).First(&mdrive) 

	// if exist create new setting drive with new day inside it
	if mquery.Error == nil {
		dayUrl, err := r.gdrive.MakeFolder(pagination.FromInt(day), mdrive.MonthUrl)
		if err != nil {
			return models.SettingDrive{}, err
		}

		settingDrive = models.SettingDrive{
			Month:    month,
			Day:      day,
			MonthUrl: mdrive.MonthUrl,
			DayUrl:   dayUrl,
			TypeID:  targetId,
		}

		resultQuery := r.db.Create(&settingDrive)
		return settingDrive, resultQuery.Error
	}

	if !errors.Is(mquery.Error, gorm.ErrRecordNotFound) {
		return models.SettingDrive{}, mquery.Error
	}

	monthUrl, err := r.gdrive.MakeFolder(month, targetId)
	if err != nil {
		return models.SettingDrive{}, err
	}

	dayUrl, err := r.gdrive.MakeFolder(pagination.FromInt(day), monthUrl)
	if err != nil {
		return models.SettingDrive{}, err
	}

	settingDrive = models.SettingDrive{
		Month:    month,
		Day:      day,
		MonthUrl: monthUrl,
		DayUrl:   dayUrl,
		TypeID:  targetId,
	}

	resultQuery := r.db.Create(&settingDrive)
	return settingDrive, resultQuery.Error
}
