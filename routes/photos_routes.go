package routes

import (
	"Deteccion_Zonas_Dengue_Backend/controllers"
	"github.com/gorilla/mux"
)

func PhotosRoute(router *mux.Router) {
	router.HandleFunc("/photo", controllers.CreatePhoto()).Methods("POST")
	router.HandleFunc("/photo/{photoId}", controllers.GetAPhoto()).Methods("GET")
	router.HandleFunc("/photo/{photoId}", controllers.EditAPhoto()).Methods("PUT")
	router.HandleFunc("/photo/{photoId}", controllers.DeleteAPhoto()).Methods("DELETE")
	router.HandleFunc("/photos", controllers.GetAllPhotos()).Methods("GET")
	router.HandleFunc("/photos", controllers.DeleteAllPhotos()).Methods("DELETE")
}
