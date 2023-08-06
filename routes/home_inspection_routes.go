package routes

import (
	"DENV_Backend/controllers"
	"github.com/go-chi/chi/v5"
)

type HomeInspectionResource struct{}

// Routes Rutas para inspección de viviendas (home inspection)
func (hir HomeInspectionResource) Routes() chi.Router {
	router := chi.NewRouter()

	// CRUD para inspección de viviendas (home inspection)
	router.Method("POST", "/", controllers.CreateHomeInspection())
	router.Method("GET", "/{homeInspectionId}", controllers.GetHomeInspection())
	router.Method("PUT", "/{homeInspectionId}", controllers.EditHomeInspection())
	router.Method("DELETE", "/{homeInspectionId}", controllers.DeleteHomeInspection())

	// Obtener inspecciones de viviendas (home inspection) por rango de fechas
	router.Method("GET", "/detailed", controllers.GetAllHomeInspectionsDetailed())
	router.Method("GET", "/summarized", controllers.GetAllHomeInspectionsSummarized())

	// Endpoint para obtener clusters de inspecciones de viviendas
	router.Method("GET", "/clusters", controllers.GetHomeInspectionClusters())

	return router
}
