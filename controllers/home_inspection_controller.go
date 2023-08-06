package controllers

import (
	"DENV_Backend/configs"
	"DENV_Backend/models"
	"DENV_Backend/responses"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"time"
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

func EditHomeInspection() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		// Obtener el ID de la inspección de vivienda
		homeInspectionID := chi.URLParam(request, "homeInspectionId")

		// Validar que el ID de la inspección de vivienda exista
		var homeInspection models.HomeInspection
		if configs.DB.First(&homeInspection, homeInspectionID).RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusNotFound,
				Message: "No se ha encontrado la inspección de vivienda",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Validar que el body está en formato JSON
		var homeInspectionData models.HomeInspection
		if err := json.NewDecoder(request.Body).Decode(&homeInspectionData); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "El cuerpo de la solicitud no está en formato JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Actualizar inspección de vivienda
		configs.DB.Model(&homeInspection).Updates(homeInspectionData)

		// Actualizar fila de dirección
		configs.DB.First(&homeInspection.Address, homeInspection.AddressID).Updates(homeInspectionData.Address)

		// Actualizar fila de tipo de contenedor
		configs.DB.First(&homeInspection.TypeContainer, homeInspection.TypeContainerID).Updates(homeInspectionData.TypeContainer)

		// Actualizar fila de tanque elevado
		configs.DB.First(&homeInspection.TypeContainer.ElevatedTank, homeInspection.TypeContainer.ElevatedTankID).Updates(homeInspectionData.TypeContainer.ElevatedTank)

		// Actualizar fila de tanque bajo
		configs.DB.First(&homeInspection.TypeContainer.LowTank, homeInspection.TypeContainer.LowTankID).Updates(homeInspectionData.TypeContainer.LowTank)

		// Actualizar fila de barriles cilindro
		configs.DB.First(&homeInspection.TypeContainer.CylinderBarrel, homeInspection.TypeContainer.CylinderBarrelID).Updates(homeInspectionData.TypeContainer.CylinderBarrel)

		// Actualizar fila de tinas baldes
		configs.DB.First(&homeInspection.TypeContainer.BucketTub, homeInspection.TypeContainer.BucketTubID).Updates(homeInspectionData.TypeContainer.BucketTub)

		// Actualizar fila de llantas
		configs.DB.First(&homeInspection.TypeContainer.Tire, homeInspection.TypeContainer.TireID).Updates(homeInspectionData.TypeContainer.Tire)

		// Actualizar fila de floreros
		configs.DB.First(&homeInspection.TypeContainer.Flower, homeInspection.TypeContainer.FlowerID).Updates(homeInspectionData.TypeContainer.Flower)

		// Actualizar fila de inservibles
		configs.DB.First(&homeInspection.TypeContainer.Useless, homeInspection.TypeContainer.UselessID).Updates(homeInspectionData.TypeContainer.Useless)

		// Actualizar fila de otros
		configs.DB.First(&homeInspection.TypeContainer.Others, homeInspection.TypeContainer.OthersID).Updates(homeInspectionData.TypeContainer.Others)

		// Actualizar fila de condiciones de vivienda
		configs.DB.First(&homeInspection.HomeCondition, homeInspection.HomeConditionID).Updates(homeInspectionData.HomeCondition)

		// Actualizar fila de total de contenedores
		configs.DB.First(&homeInspection.TotalContainer, homeInspection.TotalContainerID).Updates(homeInspectionData.TotalContainer)

		// Actualizar fila de focos de aegypti
		configs.DB.First(&homeInspection.AegyptiFocus, homeInspection.AegyptiFocusID).Updates(homeInspectionData.AegyptiFocus)

		writer.WriteHeader(http.StatusOK)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusOK,
			Message: "Inspección de vivienda actualizada con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
	}
}

func DeleteHomeInspection() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		// Obtener el ID de la inspección de vivienda
		homeInspectionID := chi.URLParam(request, "homeInspectionId")

		// Validar que el ID de la inspección de vivienda exista
		var homeInspection models.HomeInspection
		if configs.DB.First(&homeInspection, homeInspectionID).RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusNotFound,
				Message: "No se ha encontrado la inspección de vivienda",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Eliminar inspección de vivienda
		configs.DB.Delete(&homeInspection)

		// Eliminar fila de dirección
		configs.DB.Delete(&homeInspection.Address, homeInspection.AddressID)

		// Eliminar fila de tipo de contenedor
		configs.DB.Delete(&homeInspection.TypeContainer, homeInspection.TypeContainerID)

		// Eliminar fila de tanque elevado
		configs.DB.Delete(&homeInspection.TypeContainer.ElevatedTank, homeInspection.TypeContainer.ElevatedTankID)

		// Eliminar fila de tanque bajo
		configs.DB.Delete(&homeInspection.TypeContainer.LowTank, homeInspection.TypeContainer.LowTankID)

		// Eliminar fila de barriles cilindro
		configs.DB.Delete(&homeInspection.TypeContainer.CylinderBarrel, homeInspection.TypeContainer.CylinderBarrelID)

		// Eliminar fila de tinas baldes
		configs.DB.Delete(&homeInspection.TypeContainer.BucketTub, homeInspection.TypeContainer.BucketTubID)

		// Eliminar fila de llantas
		configs.DB.Delete(&homeInspection.TypeContainer.Tire, homeInspection.TypeContainer.TireID)

		// Eliminar fila de floreros
		configs.DB.Delete(&homeInspection.TypeContainer.Flower, homeInspection.TypeContainer.FlowerID)

		// Eliminar fila de inservibles
		configs.DB.Delete(&homeInspection.TypeContainer.Useless, homeInspection.TypeContainer.UselessID)

		// Eliminar fila de otros
		configs.DB.Delete(&homeInspection.TypeContainer.Others, homeInspection.TypeContainer.OthersID)

		// Eliminar fila de condiciones de vivienda
		configs.DB.Delete(&homeInspection.HomeCondition, homeInspection.HomeConditionID)

		// Eliminar fila de total de contenedores
		configs.DB.Delete(&homeInspection.TotalContainer, homeInspection.TotalContainerID)

		// Eliminar fila de focos de aegypti
		configs.DB.Delete(&homeInspection.AegyptiFocus, homeInspection.AegyptiFocusID)

		writer.WriteHeader(http.StatusOK)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusOK,
			Message: "Inspección de vivienda eliminada con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
	}
}

func GetAllHomeInspectionsSummarized() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		// Obtener fecha de inicio de parámetro de consulta
		startDate, err := time.Parse("02-01-2006", request.URL.Query().Get("startDate"))
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "La fecha de inicio debe tener el formato dd-mm-aaaa",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Obtener fecha de fin de parámetro de consulta
		endDate, err := time.Parse("02-01-2006", request.URL.Query().Get("endDate"))
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "La fecha de fin debe tener el formato dd-mm-aaaa",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Obtener inspecciones de vivienda dentro del rango de fechas y con el salto de página
		var homeInspections []models.HomeInspection
		configs.DB.Where("datetime BETWEEN ? AND ?", startDate, endDate).Find(&homeInspections).Order("datetime desc")

		// Validar que existan inspecciones de vivienda
		if len(homeInspections) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusNotFound,
				Message: "No se han encontrado inspecciones de vivienda",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Crear estructura de respuesta
		var homeInspectionsSummarized []models.HomeInspectionSummarized
		for _, homeInspection := range homeInspections {
			homeInspectionsSummarized = append(homeInspectionsSummarized, models.HomeInspectionSummarized{
				ID:        homeInspection.ID,
				Latitude:  homeInspection.Latitude,
				Longitude: homeInspection.Longitude,
				Datetime:  homeInspection.Datetime,
				PhotoUrl:  homeInspection.PhotoUrl,
			})
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusOK,
			Message: "Se han encontrado " + strconv.Itoa(len(homeInspectionsSummarized)) + " inspecciones de vivienda",
			Data:    homeInspectionsSummarized,
		}
		_ = json.NewEncoder(writer).Encode(response)
	}
}
