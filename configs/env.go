package configs

import (
	"DENV_Backend/models"
	"github.com/joho/godotenv"
	"log"
	"os"
	"runtime"
)

func GetPostgresProperties() (models.PostgresProperties, error) {
	// Detección de SO para la ruta de .env
	var err error = nil

	if runtime.GOOS == "windows" {
		err = godotenv.Load(".env")
	} else {
		err = godotenv.Load("/home/piero/DENV_Backend/.env")
	}

	if err != nil {
		log.Fatal("Error al cargar archivo .env")
		return models.PostgresProperties{}, err
	}

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
