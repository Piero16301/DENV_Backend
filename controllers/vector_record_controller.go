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

var validateVectorRecord = validator.New()

func CreateVectorRecord() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Validar que el body está en formato JSON
		var vectorRecord models.VectorRecord
		if err := json.NewDecoder(request.Body).Decode(&vectorRecord); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.VectorRecordResponse{
				Status:  http.StatusBadRequest,
				Message: "El cuerpo de la solicitud debe estar en formato JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Se valida que el body tenga los campos requeridos
		if validationErr := validateVectorRecord.Struct(vectorRecord); validationErr != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.VectorRecordResponse{
				Status:  http.StatusBadRequest,
				Message: "No se han enviado todos los campos requeridos",
				Data:    validationErr.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Insertar registro de vector
		configs.DB.Create(&vectorRecord)

		writer.WriteHeader(http.StatusCreated)
		response := responses.VectorRecordResponse{
			Status:  http.StatusCreated,
			Message: "Registro de vector creado con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
	}
}

func GetVectorRecord() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Obtener el ID del registro de vector
		vectorRecordId := chi.URLParam(request, "vectorRecordId")

		// Validar que el ID del registro de vector exista
		var vectorRecord models.VectorRecord
		if configs.DB.Preload("Address").First(&vectorRecord, vectorRecordId).RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)
			response := responses.VectorRecordResponse{
				Status:  http.StatusNotFound,
				Message: "No se ha encontrado el registro de vector",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Retornar el registro de vector
		writer.WriteHeader(http.StatusOK)
		response := responses.VectorRecordResponse{
			Status:  http.StatusOK,
			Message: "Registro de vector obtenido con éxito",
			Data:    vectorRecord,
		}
		_ = json.NewEncoder(writer).Encode(response)
	}
}

func EditVectorRecord() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Obtener el ID del registro de vector
		vectorRecordId := chi.URLParam(request, "vectorRecordId")

		// Validar que el ID del registro de vector exista
		var vectorRecord models.VectorRecord
		if configs.DB.First(&vectorRecord, vectorRecordId).RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)
			response := responses.VectorRecordResponse{
				Status:  http.StatusNotFound,
				Message: "No se ha encontrado el registro de vector",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Validar que el body está en formato JSON
		var vectorRecordData models.VectorRecord
		if err := json.NewDecoder(request.Body).Decode(&vectorRecordData); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.VectorRecordResponse{
				Status:  http.StatusBadRequest,
				Message: "El cuerpo de la solicitud debe estar en formato JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Actualizar registro de vector
		configs.DB.Model(&vectorRecord).Updates(vectorRecordData)

		// Actualizar fila de dirección
		configs.DB.First(&vectorRecord.Address, vectorRecord.AddressID).Updates(vectorRecordData.Address)

		writer.WriteHeader(http.StatusOK)
		response := responses.VectorRecordResponse{
			Status:  http.StatusOK,
			Message: "Registro de vector actualizado con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
	}
}

func DeleteVectorRecord() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Obtener el ID del registro de vector
		vectorRecordId := chi.URLParam(request, "vectorRecordId")

		// Validar que el ID del registro de vector exista
		var vectorRecord models.VectorRecord
		if configs.DB.First(&vectorRecord, vectorRecordId).RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)
			response := responses.VectorRecordResponse{
				Status:  http.StatusNotFound,
				Message: "No se ha encontrado el registro de vector",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Eliminar registro de vector
		configs.DB.Delete(&vectorRecord)

		// Eliminar fila de dirección
		configs.DB.Delete(&vectorRecord.Address, vectorRecord.AddressID)

		writer.WriteHeader(http.StatusOK)
		response := responses.VectorRecordResponse{
			Status:  http.StatusOK,
			Message: "Registro de vector eliminado con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
	}
}

func GetAllVectorRecordsDetailed() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Obtener fecha de inicio de parámetro de consulta
		startDate, err := time.Parse("02-01-2006", request.URL.Query().Get("startDate"))
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.VectorRecordResponse{
				Status:  http.StatusBadRequest,
				Message: "La fecha de inicio debe estar en formato dd-mm-aaaa",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Obtener fecha de fin de parámetro de consulta
		endDate, err := time.Parse("02-01-2006", request.URL.Query().Get("endDate"))
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.VectorRecordResponse{
				Status:  http.StatusBadRequest,
				Message: "La fecha de fin debe estar en formato dd-mm-aaaa",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Obtener todos los registros de vector
		var vectorRecords []models.VectorRecord
		configs.DB.Preload("Address").Where("datetime BETWEEN ? AND ?", startDate, endDate).Find(&vectorRecords)

		// Validar que existan registros de vector
		if len(vectorRecords) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			response := responses.VectorRecordResponse{
				Status:  http.StatusNotFound,
				Message: "No se han encontrado registros de vector",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.VectorRecordResponse{
			Status:  http.StatusOK,
			Message: "Se han encontrado " + strconv.Itoa(len(vectorRecords)) + " registros de vector",
			Data:    vectorRecords,
		}
		_ = json.NewEncoder(writer).Encode(response)
	}
}
