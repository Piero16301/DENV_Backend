package main

import (
	"Deteccion_Zonas_Dengue_Backend/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Enrutador de endpoints
	router := mux.NewRouter()

	// Rutas para reportes de casos
	routes.CaseReportRoute(router)

	// Rutas para zonas de propagación
	routes.PropagationZoneRoute(router)

	// Iniciar servidor
	log.Fatal(http.ListenAndServe(":80", router))
}
