package rds

import (
	"fmt"

	"github.com/fabioods/go-orders/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.RdsClient.Host, cfg.RdsClient.Port, cfg.RdsClient.User, cfg.RdsClient.Password, cfg.RdsClient.DBName, cfg.RdsClient.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		err := fmt.Sprintf("Error to open database connection %s", err)
		panic(err)
	}

	return db
}
