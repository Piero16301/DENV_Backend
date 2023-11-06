package configs

import (
	"DENV_Backend/models"
)

func GetPostgresProperties() (models.PostgresProperties, error) {
	pgProperties := models.PostgresProperties{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		DBName:   "denv",
		Port:     "5432",
		SSLMode:  "disable",
		TimeZone: "America/Lima",
	}

	return pgProperties, nil
}
