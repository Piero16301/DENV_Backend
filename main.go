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

	// Rutas para inspección de viviendas (home inspection)
	routes.CaseReportRoute(router)

	// Rutas para registro de vector (vector record)
	routes.PropagationZoneRoute(router)

	// Iniciar servidor en el puerto 80
	log.Fatal(http.ListenAndServe(":80", router))
}
