package routes

import (
	"DENV_Backend/controllers"
	"github.com/gorilla/mux"
)

func HomeInspectionRoute(router *mux.Router) {
	// CRUD routes
	router.HandleFunc("/home-inspection", controllers.CreateHomeInspection()).Methods("POST")
	router.HandleFunc("/home-inspection/{homeInspectionId}", controllers.GetHomeInspection()).Methods("GET")
	router.HandleFunc("/home-inspection/{homeInspectionId}", controllers.EditHomeInspection()).Methods("PUT")
	router.HandleFunc("/home-inspection/{homeInspectionId}", controllers.DeleteHomeInspection()).Methods("DELETE")
	router.HandleFunc("/home-inspections-detailed/{skip}", controllers.GetAllHomeInspectionsDetailed()).Methods("GET")
	router.HandleFunc("/home-inspections-summarized/{skip}", controllers.GetAllHomeInspectionsSummarized()).Methods("GET")
	router.HandleFunc("/home-inspections", controllers.DeleteAllHomeInspections()).Methods("DELETE")

	// Analysis routes
	router.HandleFunc("/home-inspection-clusters/{eps}/{minPoints}", controllers.GetHomeInspectionClusters()).Methods("GET")
}
