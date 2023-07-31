package main

import (
	"DENV_Backend/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	// Enrutador de endpoints
	router := chi.NewRouter()

	// Middleware para CORS
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	//router.Use(middleware.URLFormat)
	//router.Use(render.SetContentType(render.ContentTypeJSON))

	// Rutas para inspección de viviendas (home inspection)
	router.Mount("/home-inspections", routes.HomeInspectionResource{}.Routes())

	// Iniciar servidor en el puerto 80
	_ = http.ListenAndServe(":5000", router)
}
