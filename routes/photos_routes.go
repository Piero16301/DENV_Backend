package routes

import (
	"Deteccion_Zonas_Dengue_Backend/controllers"
	"github.com/gorilla/mux"
)

func PhotosRoute(router *mux.Router) {
	router.HandleFunc("/photo", controllers.CreatePhoto()).Methods("POST")
}
