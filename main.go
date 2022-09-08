package main

import (
	"DENV_Backend/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Enrutador de endpoints
	router := mux.NewRouter()

	// Rutas para inspección de viviendas (home inspection)
	routes.HomeInspectionRoute(router)

	// Rutas para registro de vector (vector record)
	routes.VectorRecordRoute(router)

	// Iniciar servidor en el puerto 80
	log.Fatal(http.ListenAndServe(":80", router))
}
