package routes

import (
	"Deteccion_Zonas_Dengue_Backend/controllers"
	"github.com/gorilla/mux"
)

func VectorRecordRoute(router *mux.Router) {
	router.HandleFunc("/vector-record", controllers.CreatePropagationZone()).Methods("POST")
	router.HandleFunc("vector-record/{vectorRecordId}", controllers.GetPropagationZone()).Methods("GET")
	router.HandleFunc("vector-record/{vectorRecordId}", controllers.EditPropagationZone()).Methods("PUT")
	router.HandleFunc("vector-record/{vectorRecordId}", controllers.DeletePropagationZone()).Methods("DELETE")
	router.HandleFunc("vector-records-detailed", controllers.GetAllPropagationZonesDetailed()).Methods("GET")
	router.HandleFunc("vector-records-summarized", controllers.GetAllPropagationZonesSummarized()).Methods("GET")
	router.HandleFunc("vector-records", controllers.DeleteAllPropagationZones()).Methods("DELETE")
}
