package routes

import (
	"Deteccion_Zonas_Dengue_Backend/controllers"
	"github.com/gorilla/mux"
)

func PropagationZoneRoute(router *mux.Router) {
	router.HandleFunc("/propagation-zone", controllers.CreatePropagationZone()).Methods("POST")
	router.HandleFunc("/propagation-zone/{propagationZoneId}", controllers.GetPropagationZone()).Methods("GET")
	router.HandleFunc("/propagation-zone/{propagationZoneId}", controllers.EditPropagationZone()).Methods("PUT")
	router.HandleFunc("/propagation-zone/{propagationZoneId}", controllers.DeletePropagationZone()).Methods("DELETE")
	router.HandleFunc("/propagation-zones-detailed", controllers.GetAllPropagationZonesDetailed()).Methods("GET")
	router.HandleFunc("/propagation-zones-summarized", controllers.GetAllPropagationZonesSummarized()).Methods("GET")
	router.HandleFunc("/propagation-zones", controllers.DeleteAllPropagationZones()).Methods("DELETE")
}
