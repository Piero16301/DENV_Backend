package main

import (
	"Deteccion_Zonas_Dengue_Backend/configs"
	"Deteccion_Zonas_Dengue_Backend/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// Establecer conexión con MongoDB
	configs.ConnectDB()

	// Rutas de usuario
	routes.UserRoute(router)

	// Rutas de puntos mosquitos
	routes.PointsRoute(router)

	// Rutas de fotos mosquitos
	routes.PhotosRoute(router)

	log.Fatal(http.ListenAndServe(":80", router))
}
