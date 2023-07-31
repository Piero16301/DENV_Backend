package controllers

import (
	"DENV_Backend/configs"
	"DENV_Backend/models"
	"DENV_Backend/responses"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validateHomeInspection = validator.New()

func CreateHomeInspection() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Validar que el body está en formato JSON
		var homeInspection models.HomeInspection
		if err := json.NewDecoder(request.Body).Decode(&homeInspection); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "El cuerpo de la solicitud no está en formato JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Se valida que se envíen todos los campos requeridos
		if validationErr := validateHomeInspection.Struct(&homeInspection); validationErr != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "No se han enviado todos los campos requeridos",
				Data:    validationErr.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Insertar inspección de vivienda
		configs.DB.Create(&homeInspection)

		writer.WriteHeader(http.StatusCreated)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusCreated,
			Message: "Inspección de vivienda creada con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
	}
}

func GetHomeInspection() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		// Obtener el ID de la inspección de vivienda
		homeInspectionID := chi.URLParam(request, "homeInspectionId")

		// Validar que el ID de la inspección de vivienda exista
		var homeInspection models.HomeInspection
		if configs.DB.Preload("Address").Preload("TypeContainer").Preload("HomeCondition").Preload("TotalContainer").Preload("AegyptiFocus").First(&homeInspection, homeInspectionID).RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusNotFound,
				Message: "No se ha encontrado la inspección de vivienda",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		var typeContainer models.TypeContainer
		configs.DB.Preload("ElevatedTank").Preload("LowTank").Preload("CylinderBarrel").Preload("BucketTub").Preload("Tire").Preload("Flower").Preload("Useless").Preload("Others").First(&typeContainer, homeInspection.TypeContainerID)
		homeInspection.TypeContainer = typeContainer

		// Retornar la inspección de vivienda
		writer.WriteHeader(http.StatusOK)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusOK,
			Message: "Inspección de vivienda obtenida con éxito",
			Data:    homeInspection,
		}
		_ = json.NewEncoder(writer).Encode(response)
	}
}
