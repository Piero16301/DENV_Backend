package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"runtime"
)

func GetMongoURI() string {
	// Detección de SO para la ruta de .env
	var err error = nil

	if runtime.GOOS == "windows" {
		err = godotenv.Load(".env")
	} else {
		err = godotenv.Load("/home/ubuntu/DENV_Backend/.env")
	}

	if err != nil {
		log.Fatal("Error al cargar archivo .env")
	}

	return os.Getenv("MONGOURI")
}
