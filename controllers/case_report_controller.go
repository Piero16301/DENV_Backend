package controllers

import (
	"Deteccion_Zonas_Dengue_Backend/configs"
	"Deteccion_Zonas_Dengue_Backend/models"
	"Deteccion_Zonas_Dengue_Backend/responses"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var caseReportCollection = configs.GetCollection(configs.DB, "CaseReport")
var validateCaseReport = validator.New()

func CreateCaseReport() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var caseReport models.CaseReport
		defer cancel()

		// Validar que el body está en formato JSON
		if err := json.NewDecoder(request.Body).Decode(&caseReport); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.CaseReportResponse{
				Status:  http.StatusBadRequest,
				Message: "El cuerpo de la solicitud no está en formato JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Se valida que se envíen todos los campos requeridos
		if validationErr := validateCaseReport.Struct(&caseReport); validationErr != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.CaseReportResponse{
				Status:  http.StatusBadRequest,
				Message: "No se han enviado todos los campos requeridos",
				Data:    validationErr.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Se crea instancia del reporte de caso
		newCaseReport := models.CaseReport{
			Id:        primitive.NewObjectID(),
			Address:   caseReport.Address,
			Comment:   caseReport.Comment,
			Datetime:  caseReport.Datetime,
			Latitude:  caseReport.Latitude,
			Longitude: caseReport.Longitude,
			PhotoURL:  caseReport.PhotoURL,
		}

		// Se inserta el reporte de caso en MongoDB
		_, err := caseReportCollection.InsertOne(ctx, newCaseReport)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.CaseReportResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al crear el reporte de caso",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusCreated)
		response := responses.CaseReportResponse{
			Status:  http.StatusCreated,
			Message: "Reporte de caso creado con éxito",
			Data:    newCaseReport,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Reporte de caso %s creado con éxito\n", newCaseReport.Id.Hex())
	}
}

func GetCaseReport() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(request)
		caseReportId := params["caseReportId"]
		var caseReport models.CaseReport
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(caseReportId)

		err := caseReportCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&caseReport)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.CaseReportResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al obtener el reporte de caso",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.CaseReportResponse{
			Status:  http.StatusOK,
			Message: "Reporte de caso obtenido con éxito",
			Data:    caseReport,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Reporte de caso %s leído con éxito\n", caseReportId)
	}
}

func EditCaseReport() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(request)
		caseReportId := params["caseReportId"]
		var caseReport models.CaseReport
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(caseReportId)

		// Validar que el body está en formato JSON
		if err := json.NewDecoder(request.Body).Decode(&caseReport); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.CaseReportResponse{
				Status:  http.StatusBadRequest,
				Message: "El cuerpo de la solicitud no está en formato JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Se valida que se envíen todos los campos requeridos
		if validationErr := validateCaseReport.Struct(&caseReport); validationErr != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.CaseReportResponse{
				Status:  http.StatusBadRequest,
				Message: "No se han enviado todos los campos requeridos",
				Data:    validationErr.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		bsonCaseReport := bson.M{
			"address":   caseReport.Address,
			"comment":   caseReport.Comment,
			"datetime":  caseReport.Datetime,
			"latitude":  caseReport.Latitude,
			"longitude": caseReport.Longitude,
			"photourl":  caseReport.PhotoURL,
		}

		result, err := caseReportCollection.UpdateOne(ctx, bson.M{"id": objectId}, bson.M{"$set": bsonCaseReport})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.CaseReportResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al actualizar el reporte de caso",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Actualizar los campos del reporte de caso
		var updatedCaseReport models.CaseReport
		if result.MatchedCount == 1 {
			err := caseReportCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&updatedCaseReport)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.CaseReportResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al actualizar el reporte de caso",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
				return
			}
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.CaseReportResponse{
			Status:  http.StatusOK,
			Message: "Reporte de caso actualizado con éxito",
			Data:    updatedCaseReport,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Reporte de caso %s actualizado con éxito\n", caseReportId)
	}
}

func DeleteCaseReport() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(request)
		caseReportId := params["caseReportId"]
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(caseReportId)

		result, err := caseReportCollection.DeleteOne(ctx, bson.M{"id": objectId})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.CaseReportResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al eliminar el reporte de caso",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		if result.DeletedCount < 1 {
			writer.WriteHeader(http.StatusNotFound)
			response := responses.CaseReportResponse{
				Status:  http.StatusNotFound,
				Message: "No se encontró el reporte de caso con el ID especificado",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.CaseReportResponse{
			Status:  http.StatusOK,
			Message: "Reporte de caso eliminado con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Reporte de caso %s eliminado con éxito\n", caseReportId)
	}
}

func GetAllCaseReports() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var caseReports []models.CaseReport
		defer cancel()

		results, err := caseReportCollection.Find(ctx, bson.M{})

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.CaseReportResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al obtener los reportes de casos",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Lectura de resultados de la base de datos
		defer func(results *mongo.Cursor, ctx context.Context) {
			_ = results.Close(ctx)
		}(results, ctx)

		for results.Next(ctx) {
			var singleCaseReport models.CaseReport
			if err = results.Decode(&singleCaseReport); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.CaseReportResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al obtener los reportes de casos",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
			}
			caseReports = append(caseReports, singleCaseReport)
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.CaseReportResponse{
			Status:  http.StatusOK,
			Message: "Reportes de casos obtenidos con éxito",
			Data:    caseReports,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Println("Reportes de casos leídos con éxito")
	}
}

func DeleteAllCaseReports() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err := caseReportCollection.DeleteMany(ctx, bson.M{})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.CaseReportResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al eliminar todos los reportes de casos",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.CaseReportResponse{
			Status:  http.StatusOK,
			Message: "Todos los reportes de casos eliminados con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Println("Reportes de casos eliminados con éxito")
	}
}
