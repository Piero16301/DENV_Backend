package routes

import (
	"Deteccion_Zonas_Dengue_Backend/controllers"
	"github.com/gorilla/mux"
)

func CaseReportRoute(router *mux.Router) {
	router.HandleFunc("/case-report", controllers.CreateCaseReport()).Methods("POST")
	router.HandleFunc("/case-report/{caseReportId}", controllers.GetCaseReport()).Methods("GET")
	router.HandleFunc("/case-report/{caseReportId}", controllers.EditCaseReport()).Methods("PUT")
	router.HandleFunc("/case-report/{caseReportId}", controllers.DeleteCaseReport()).Methods("DELETE")
	router.HandleFunc("/case-reports", controllers.GetAllCaseReports()).Methods("GET")
	router.HandleFunc("/case-reports", controllers.DeleteAllCaseReports()).Methods("DELETE")
}
