package controllers

import (
	"DENV_Backend/configs"
	"DENV_Backend/models"
	"DENV_Backend/responses"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

var homeInspectionCollection = configs.GetCollection(configs.DB, "HomeInspection")
var validateHomeInspection = validator.New()

func CreateHomeInspection() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var homeInspection models.HomeInspection
		defer cancel()

		// Validar que el body está en formato JSON
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

		// Se crea instancia de la inspección de vivienda
		newHomeInspection := models.HomeInspection{
			Id:                primitive.NewObjectID(),
			Address:           homeInspection.Address,
			Comment:           homeInspection.Comment,
			Datetime:          homeInspection.Datetime,
			DNI:               homeInspection.DNI,
			Latitude:          homeInspection.Latitude,
			Longitude:         homeInspection.Longitude,
			PhotoURL:          homeInspection.PhotoURL,
			NumberInhabitants: homeInspection.NumberInhabitants,
			HomeCondition:     homeInspection.HomeCondition,
			TypeContainers:    homeInspection.TypeContainers,
			TotalContainer:    homeInspection.TotalContainer,
			AegyptiFocus:      homeInspection.AegyptiFocus,
		}

		// Se inserta el reporte de caso en MongoDB
		_, err := homeInspectionCollection.InsertOne(ctx, newHomeInspection)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al crear la inspección de vivienda",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusCreated)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusCreated,
			Message: "Inspección de vivienda creada con éxito",
			Data:    newHomeInspection,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Inspección de vivienda %s creada con éxito\n", newHomeInspection.Id.Hex())
	}
}

func GetHomeInspection() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(request)
		homeInspectionId := params["homeInspectionId"]
		var homeInspection models.HomeInspection
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(homeInspectionId)

		err := homeInspectionCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&homeInspection)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al obtener la inspección de vivienda",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusOK,
			Message: "Inspección de vivienda obtenida con éxito",
			Data:    homeInspection,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Inspección de vivienda %s obtenida con éxito\n", homeInspection.Id.Hex())
	}
}

func EditHomeInspection() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(request)
		homeInspectionId := params["homeInspectionId"]
		var homeInspection models.HomeInspection
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(homeInspectionId)

		// Validar que el body está en formato JSON
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

		bsonHomeInspection := bson.M{
			"address":   homeInspection.Address,
			"comment":   homeInspection.Comment,
			"datetime":  homeInspection.Datetime,
			"dni":       homeInspection.DNI,
			"latitude":  homeInspection.Latitude,
			"longitude": homeInspection.Longitude,
			"photourl":  homeInspection.PhotoURL,

			"numberinhabitants": homeInspection.NumberInhabitants,
			"homecondition":     homeInspection.HomeCondition,
			"typecontainers":    homeInspection.TypeContainers,
			"totalcontainer":    homeInspection.TotalContainer,
			"aegyptifocus":      homeInspection.AegyptiFocus,
		}

		result, err := homeInspectionCollection.UpdateOne(ctx, bson.M{"id": objectId}, bson.M{"$set": bsonHomeInspection})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al actualizar la inspección de vivienda",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Actualizar los campos del reporte de caso
		var updatedHomeInspection models.HomeInspection
		if result.MatchedCount == 1 {
			err := homeInspectionCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&updatedHomeInspection)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.HomeInspectionResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al actualizar la inspección de vivienda",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
				return
			}
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusOK,
			Message: "Inspección de vivienda actualizada con éxito",
			Data:    updatedHomeInspection,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Inspección de vivienda %s actualizada con éxito\n", homeInspectionId)
	}
}

func DeleteHomeInspection() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(request)
		homeInspectionId := params["homeInspectionId"]
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(homeInspectionId)

		result, err := homeInspectionCollection.DeleteOne(ctx, bson.M{"id": objectId})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al eliminar la inspección de vivienda",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		if result.DeletedCount < 1 {
			writer.WriteHeader(http.StatusNotFound)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusNotFound,
				Message: "No se encontró la inspección de vivienda con el ID especificado",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusOK,
			Message: "Inspección de vivienda eliminada con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Inspección de vivienda %s eliminada con éxito", homeInspectionId)
	}
}

func GetAllHomeInspectionsDetailed() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var homeInspections []models.HomeInspection
		defer cancel()

		results, err := homeInspectionCollection.Find(ctx, bson.M{})

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al obtener las inspecciones de vivienda",
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
			var singleHomeInspection models.HomeInspection
			if err = results.Decode(&singleHomeInspection); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.HomeInspectionResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al obtener las inspecciones de vivienda",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
			}
			homeInspections = append(homeInspections, singleHomeInspection)
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusOK,
			Message: "Inspecciones de vivienda detalladas obtenidas con éxito",
			Data:    homeInspections,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Println("Inspecciones de vivienda detalladas obtenidas con éxito")
	}
}

func GetAllHomeInspectionsSummarized() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var homeInspections []models.HomeInspection
		defer cancel()

		results, err := homeInspectionCollection.Find(ctx, bson.M{}, &options.FindOptions{Projection: bson.M{
			"id":        1,
			"latitude":  1,
			"longitude": 1,
		}})

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al obtener las inspecciones de vivienda",
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
			var singleHomeInspection models.HomeInspection
			if err = results.Decode(&singleHomeInspection); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.HomeInspectionResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al obtener las inspecciones de vivienda resumidas",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
			}
			homeInspections = append(homeInspections, singleHomeInspection)
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusOK,
			Message: "Inspecciones de vivienda resumidas obtenidas con éxito",
			Data:    homeInspections,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Println("Inspecciones de vivienda resumidas obtenidas con éxito")
	}
}

func DeleteAllHomeInspections() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err := homeInspectionCollection.DeleteMany(ctx, bson.M{})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al eliminar todas las inspecciones de vivienda",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusOK,
			Message: "Todos las inspecciones de vivienda eliminadas con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Println("Inspecciones de vivienda eliminadas con éxito")
	}
}
