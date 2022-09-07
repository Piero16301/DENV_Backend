package routes

import (
	"Deteccion_Zonas_Dengue_Backend/controllers"
	"github.com/gorilla/mux"
)

func HomeInspectionRoute(router *mux.Router) {
	router.HandleFunc("/home-inspection", controllers.CreateHomeInspection()).Methods("POST")
	router.HandleFunc("/home-inspection/{homeInspectionId}", controllers.GetHomeInspection()).Methods("GET")
	router.HandleFunc("/home-inspection/{homeInspectionId}", controllers.EditHomeInspection()).Methods("PUT")
	router.HandleFunc("/home-inspection/{homeInspectionId}", controllers.DeleteHomeInspection()).Methods("DELETE")
	router.HandleFunc("/home-inspections-detailed", controllers.GetAllHomeInspectionsDetailed()).Methods("GET")
	router.HandleFunc("/home-inspections-summarized", controllers.GetAllHomeInspectionsSummarized()).Methods("GET")
	router.HandleFunc("/home-inspections", controllers.DeleteAllHomeInspections()).Methods("DELETE")
}
