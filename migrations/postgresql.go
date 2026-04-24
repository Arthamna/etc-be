package database

import (
	"arthamna/rplLibrary/internal/models"
	// "arthamna/rplLibrary/pkg/database/migrations"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToPostgresql() *gorm.DB {
    dbUSER := os.Getenv("DB_USER")
    dbPASSWORD := os.Getenv("DB_PASS")
    dbHOST := os.Getenv("DB_HOST")
    dbPORT := os.Getenv("DB_PORT")
    dbDBNAME := os.Getenv("DB_DBNAME")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=prefer",
        dbHOST, dbUSER, dbPASSWORD, dbDBNAME, dbPORT)

    newLogger := logger.New(
        log.New(log.Writer(), "\r\n", log.LstdFlags),
        logger.Config{
            SlowThreshold: time.Second,
            LogLevel:      logger.Info,
            Colorful:      true,
        },
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: newLogger,
        DisableForeignKeyConstraintWhenMigrating: true,
    })
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }

    fmt.Println("Connected to PostgreSQL successfully!")

    if err := AutoMigrateAll(db); err != nil {
        log.Fatal(err)
    }

        return db
}

func AutoMigrateAll(db *gorm.DB) error {
    // refresh table
	// db.Migrator().DropConstraint(&models.Article{}, "fk_categories_articles")
	// db.Migrator().DropConstraint(&models.Article{}, "fk_articles_author")
	// db.Migrator().DropTable(&models.Category{}, &models.User{} , &models.Article{}) 
	
	if err := db.AutoMigrate(
        &models.User{}, 
        &models.Category{}, 
        &models.Book{},
        &models.BookBorrowing{},
        &models.BookCategory{},
    ); err != nil {
		return err
	}
	
	return nil
}