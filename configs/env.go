package configs

import (
	"DENV_Backend/models"
	"os"
)

func GetPostgresProperties() (models.PostgresProperties, error) {
	pgProperties := models.PostgresProperties{
		Host:     os.Getenv("PGHOST"),
		User:     os.Getenv("PGUSER"),
		Password: os.Getenv("PGPASSWORD"),
		DBName:   os.Getenv("PGDATABASE"),
		Port:     os.Getenv("PGPORT"),
		SSLMode:  os.Getenv("PGSSLMODE"),
		TimeZone: os.Getenv("PGTZ"),
	}

	return pgProperties, nil
}
