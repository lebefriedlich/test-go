package db

import (
	"my-go-project/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
