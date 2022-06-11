package main

import (
	"Deteccion_Zonas_Dengue_Backend/configs"
	"Deteccion_Zonas_Dengue_Backend/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Enrutador de endpoints
	router := mux.NewRouter()

	// Establecer conexión con MongoDB
	configs.ConnectDB()

	// Rutas para reportes de casos y zonas de propagación
	routes.CaseReportRoute(router)
	routes.PropagationZoneRoute(router)

	// Iniciar servidor
	log.Fatal(http.ListenAndServe(":80", router))
}
