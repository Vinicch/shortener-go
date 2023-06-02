package repository

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createPgConnection() *gorm.DB {
	url := os.Getenv("DATABASE_URL")
	if strings.TrimSpace(url) == "" {
		log.Fatal().Msg("DATABASE_URL not configured in environment")
	}

	db, err := gorm.Open(postgres.Open(url))
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to the database")
	}

	log.Debug().Msg("Database connection established")

	return db
}
