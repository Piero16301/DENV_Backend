package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"runtime"
)

func EnvMongoURI() string {
	// Detección de SO para la ruta de .env
	var err error = nil
	if runtime.GOOS == "windows" {
		err = godotenv.Load(".env")
	} else {
		err = godotenv.Load("/home/piero/Golang_Backend/REST_API_MongoDB/.env")
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}