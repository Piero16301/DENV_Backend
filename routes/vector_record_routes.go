package routes

import (
	"DENV_Backend/controllers"
	"github.com/go-chi/chi/v5"
)

type VectorRecordResource struct{}

// Routes Rutas para registro de vectores (vector record)
func (vrr VectorRecordResource) Routes() chi.Router {
	router := chi.NewRouter()

	// CRUD para registros de vectores (vector record)
	router.Method("POST", "/", controllers.CreateVectorRecord())
	router.Method("GET", "/{vectorRecordId}", controllers.GetVectorRecord())
	router.Method("PUT", "/{vectorRecordId}", controllers.EditVectorRecord())

	return router
}
