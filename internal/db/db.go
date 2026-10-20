package db

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(ctx context.Context) (*gorm.DB, error) {
	host := os.Getenv("DB_HOST_URL")
	dbName := os.Getenv("DB_DATABASE_NAME")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	sslMode := os.Getenv("DB_SSLMODE")

	dsn := (&url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(username, password),
		Host:   host,
		Path:   dbName,
		RawQuery: url.Values{
			"sslmode": []string{sslMode},
		}.Encode(),
	}).String()

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableAutomaticPing: true,
	})
	if err != nil {
		return nil, fmt.Errorf("open PostgreSQL database: %w", err)
	}

	sqlDB, err := database.DB()
	if err != nil {
		return nil, fmt.Errorf("get PostgreSQL connection pool: %w", err)
	}
	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("ping PostgreSQL database: %w", err)
	}

	database.AutoMigrate(&Event{})

	return database, nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.Close()
	}
	return err
}
