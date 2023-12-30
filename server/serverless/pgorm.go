package pgorm

import (
	"fmt"
	"log"

	utils "server/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitOrm(config *utils.PGConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host,
		config.Username,
		config.Password,
		config.Dbname,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}
