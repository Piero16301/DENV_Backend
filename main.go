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

	// Rutas para reporte de casos
	routes.CaseReportRoute(router)

	// Rutas para zona de propagación
	routes.PropagationZoneRoute(router)

	log.Fatal(http.ListenAndServe(":80", router))
}
