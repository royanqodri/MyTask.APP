package database

import (
	"fmt"
	"royan/cleanarch/app/config"

	"gorm.io/driver/postgres" // Import PostgreSQL driver
	"gorm.io/gorm"
)

func InitDBPostgreSQL(cfg *config.AppConfig) *gorm.DB {
	// Update the connection string for PostgreSQL
	// user=username password=password dbname=db_name host=hostdb port=portdb sslmode=disable
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_HOSTNAME, cfg.DB_PORT)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
