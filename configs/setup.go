package configs

import (
	"DENV_Backend/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() *gorm.DB {
	fmt.Println("Conectando a base de datos...")

	pgProperties, err := GetPostgresProperties()
	if err != nil {
		log.Fatal(err)
	}

	client, err := gorm.Open(postgres.Open(pgProperties.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado a base de datos <" + pgProperties.DBName + "> en " + pgProperties.Host + ":" + pgProperties.Port)

	_ = client.AutoMigrate(
		&models.HomeInspection{},
		&models.VectorRecord{},
		&models.Address{},
		&models.Container{},
		&models.TypeContainer{},
		&models.HomeCondition{},
		&models.TotalContainer{},
		&models.AegyptiFocus{},
	)

	return client
}

// DB Instancia de Cliente
var DB = ConnectDB()
