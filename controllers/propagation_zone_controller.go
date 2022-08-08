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
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

var propagationZoneCollection = configs.GetCollection(configs.DB, "PropagationZone")
var validatePropagationZone = validator.New()

func CreatePropagationZone() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var propagationZone models.PropagationZone
		defer cancel()

		// Validar que el body está en formato JSON
		if err := json.NewDecoder(request.Body).Decode(&propagationZone); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.PropagationZoneResponse{
				Status:  http.StatusBadRequest,
				Message: "El cuerpo de la solicitud no está en formato JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Se valida que se envíen todos los campos requeridos
		if validationErr := validatePropagationZone.Struct(&propagationZone); validationErr != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.PropagationZoneResponse{
				Status:  http.StatusBadRequest,
				Message: "No se han enviado todos los campos requeridos",
				Data:    validationErr.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Se crea instancia de zona de propagación
		newPropagationZone := models.PropagationZone{
			Id:        primitive.NewObjectID(),
			Address:   propagationZone.Address,
			Comment:   propagationZone.Comment,
			Datetime:  propagationZone.Datetime,
			Latitude:  propagationZone.Latitude,
			Longitude: propagationZone.Longitude,
			PhotoURL:  propagationZone.PhotoURL,
		}

		// Se inserta la zona de propagación en MongoDB
		_, err := propagationZoneCollection.InsertOne(ctx, newPropagationZone)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.PropagationZoneResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al crear la zona de propagación",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusCreated)
		response := responses.PropagationZoneResponse{
			Status:  http.StatusCreated,
			Message: "Zona de propagación creada con éxito",
			Data:    newPropagationZone,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Zona de propagación %s creada con éxito\n", newPropagationZone.Id.Hex())
	}
}

func GetPropagationZone() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(request)
		propagationZoneId := params["propagationZoneId"]
		var propagationZone models.PropagationZone
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(propagationZoneId)

		err := propagationZoneCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&propagationZone)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.PropagationZoneResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al obtener la zona de propagación",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.PropagationZoneResponse{
			Status:  http.StatusOK,
			Message: "Zona de propagación obtenida con éxito",
			Data:    propagationZone,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Zona de propagación %s obtenida con éxito\n", propagationZoneId)
	}
}

func EditPropagationZone() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(request)
		propagationZoneId := params["propagationZoneId"]
		var propagationZone models.PropagationZone
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(propagationZoneId)

		// Validar que el body está en formato JSON
		if err := json.NewDecoder(request.Body).Decode(&propagationZone); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.PropagationZoneResponse{
				Status:  http.StatusBadRequest,
				Message: "El cuerpo de la solicitud no está en formato JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Se valida que se envíen todos los campos requeridos
		if validationErr := validatePropagationZone.Struct(&propagationZone); validationErr != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.PropagationZoneResponse{
				Status:  http.StatusBadRequest,
				Message: "No se han enviado todos los campos requeridos",
				Data:    validationErr.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		bsonPropagationZone := bson.M{
			"address":   propagationZone.Address,
			"comment":   propagationZone.Comment,
			"datetime":  propagationZone.Datetime,
			"latitude":  propagationZone.Latitude,
			"longitude": propagationZone.Longitude,
			"photourl":  propagationZone.PhotoURL,
		}

		result, err := propagationZoneCollection.UpdateOne(ctx, bson.M{"id": objectId}, bson.M{"$set": bsonPropagationZone})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.PropagationZoneResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al actualizar la zona de propagación",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Actualizar los campos de la zona de propagación
		var updatedPropagationZone models.PropagationZone
		if result.MatchedCount == 1 {
			err := propagationZoneCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&updatedPropagationZone)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.PropagationZoneResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al actualizar la zona de propagación",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
				return
			}
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.PropagationZoneResponse{
			Status:  http.StatusOK,
			Message: "Zona de propagación actualizada con éxito",
			Data:    updatedPropagationZone,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Zona de propagación %s actualizada con éxito\n", propagationZoneId)
	}
}

func DeletePropagationZone() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(request)
		propagationZoneId := params["propagationZoneId"]
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(propagationZoneId)

		result, err := propagationZoneCollection.DeleteOne(ctx, bson.M{"id": objectId})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.PropagationZoneResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al eliminar la zona de propagación",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		if result.DeletedCount < 1 {
			writer.WriteHeader(http.StatusNotFound)
			response := responses.PropagationZoneResponse{
				Status:  http.StatusNotFound,
				Message: "No se encontró la zona de propagación con el ID especificado",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.PropagationZoneResponse{
			Status:  http.StatusOK,
			Message: "Zona de propagación eliminada con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Zona de propagación %s eliminada con éxito\n", propagationZoneId)
	}
}

func GetAllPropagationZonesDetailed() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var propagationZones []models.PropagationZone
		defer cancel()

		results, err := propagationZoneCollection.Find(ctx, bson.M{})

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.PropagationZoneResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al obtener las zonas de propagación detalladas",
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
			var singlePropagationZone models.PropagationZone
			if err = results.Decode(&singlePropagationZone); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.PropagationZoneResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al obtener las zonas de propagación detalladas",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
			}
			propagationZones = append(propagationZones, singlePropagationZone)
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.PropagationZoneResponse{
			Status:  http.StatusOK,
			Message: "Zonas de propagación detalladas obtenidas con éxito",
			Data:    propagationZones,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Println("Zonas de propagación detalladas obtenidas con éxito")
	}
}

func GetAllPropagationZonesSummarized() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var propagationZones []models.PropagationZone
		defer cancel()

		results, err := propagationZoneCollection.Find(ctx, bson.M{}, &options.FindOptions{Projection: bson.M{
			"id":        1,
			"latitude":  1,
			"longitude": 1,
		}})

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.PropagationZoneResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al obtener las zonas de propagación resumidas",
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
			var singlePropagationZone models.PropagationZone
			if err = results.Decode(&singlePropagationZone); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.PropagationZoneResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al obtener las zonas de propagación resumidas",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
			}
			propagationZones = append(propagationZones, singlePropagationZone)
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.PropagationZoneResponse{
			Status:  http.StatusOK,
			Message: "Zonas de propagación resumidas obtenidas con éxito",
			Data:    propagationZones,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Println("Zonas de propagación resumidas obtenidas con éxito")
	}
}

func DeleteAllPropagationZones() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err := propagationZoneCollection.DeleteMany(ctx, bson.M{})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.PropagationZoneResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al eliminar todas las zonas de propagación",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.PropagationZoneResponse{
			Status:  http.StatusOK,
			Message: "Todas las zonas de propagación eliminadas con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Println("Zonas de propagación eliminadas con éxito")
	}
}
