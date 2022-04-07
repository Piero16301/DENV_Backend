package routes

import (
	"Deteccion_Zonas_Dengue_Backend/controllers"
	"github.com/gorilla/mux"
)

func PointsRoute(router *mux.Router) {
	router.HandleFunc("/point", controllers.CreatePoint()).Methods("POST")
	router.HandleFunc("/point/{pointId}", controllers.GetAPoint()).Methods("GET")
}
