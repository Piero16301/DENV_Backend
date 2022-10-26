package routes

import (
	"DENV_Backend/controllers"
	"github.com/gorilla/mux"
)

func VectorRecordRoute(router *mux.Router) {
	// CRUD routes
	router.HandleFunc("/vector-record", controllers.CreateVectorRecord()).Methods("POST")
	router.HandleFunc("/vector-record/{vectorRecordId}", controllers.GetVectorRecord()).Methods("GET")
	router.HandleFunc("/vector-record/{vectorRecordId}", controllers.EditVectorRecord()).Methods("PUT")
	router.HandleFunc("/vector-record/{vectorRecordId}", controllers.DeleteVectorRecord()).Methods("DELETE")
	router.HandleFunc("/vector-records-detailed/{skip}", controllers.GetAllVectorRecordsDetailed()).Methods("GET")
	router.HandleFunc("/vector-records-summarized/{skip}", controllers.GetAllVectorRecordsSummarized()).Methods("GET")
	router.HandleFunc("/vector-records", controllers.DeleteAllVectorRecords()).Methods("DELETE")

	// Analysis routes
	router.HandleFunc("/vector-record-clusters/{eps}/{minPoints}", controllers.GetVectorRecordClusters()).Methods("GET")
}
