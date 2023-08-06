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
	//router.Method("GET", "/detailed/{skip}", controllers.GetAllHomeInspectionsDetailed())
	//router.Method("GET", "/summarized/{skip}", controllers.GetAllHomeInspectionsSummarized())
	//router.Method("DELETE", "/", controllers.DeleteAllHomeInspections())
	//
	//// Endpoint para obtener clusters de inspecciones de viviendas
	//router.Method("GET", "/clusters/{eps}/{minPoints}", controllers.GetHomeInspectionClusters())

	return router
}
