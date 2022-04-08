package routes

import (
	"Deteccion_Zonas_Dengue_Backend/controllers"
	"github.com/gorilla/mux"
)

func PointsRoute(router *mux.Router) {
	router.HandleFunc("/point", controllers.CreatePoint()).Methods("POST")
	router.HandleFunc("/point/{pointId}", controllers.GetAPoint()).Methods("GET")
	router.HandleFunc("/point/{pointId}", controllers.EditAPoint()).Methods("PUT")
	router.HandleFunc("/point/{pointId}", controllers.DeleteAPoint()).Methods("DELETE")
	router.HandleFunc("/points", controllers.GetAllPoints()).Methods("GET")
	router.HandleFunc("/points", controllers.DeleteAllPoints()).Methods("DELETE")
}
