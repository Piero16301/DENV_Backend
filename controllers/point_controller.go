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
	"net/http"
	"time"
)

var pointCollection = configs.GetCollection(configs.DB, "points")
var validatePoint = validator.New()

func CreatePoint() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var point models.Point
		defer cancel()

		// Validar el body del request
		if err := json.NewDecoder(r.Body).Decode(&point); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.PointResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			_ = json.NewEncoder(rw).Encode(response)
			return
		}

		// Usar la librería para validar los campos requeridos
		if validationErr := validatePoint.Struct(&point); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.PointResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			_ = json.NewEncoder(rw).Encode(response)
			return
		}

		newPoint := models.Point{
			Id:        primitive.NewObjectID(),
			Comment:   point.Comment,
			Address:   point.Address,
			DateTime:  point.DateTime,
			PhotoURL:  point.PhotoURL,
			Latitude:  point.Latitude,
			Longitude: point.Longitude,
		}
		result, err := pointCollection.InsertOne(ctx, newPoint)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.PointResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			_ = json.NewEncoder(rw).Encode(response)
			return
		}
		rw.WriteHeader(http.StatusCreated)
		response := responses.PointResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		_ = json.NewEncoder(rw).Encode(response)
		fmt.Println("Nuevo usuario creado con éxito")
	}
}

func GetAPoint() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		pointId := params["pointId"]
		var point models.Point
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(pointId)

		err := pointCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&point)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.PointResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			_ = json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.PointResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": point}}
		_ = json.NewEncoder(rw).Encode(response)
		fmt.Printf("Punto %s leído con éxito\n", pointId)
	}
}

func EditAPoint() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		pointId := params["pointId"]
		var point models.Point
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(pointId)

		// Validar el body del request
		if err := json.NewDecoder(r.Body).Decode(&point); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.PointResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			_ = json.NewEncoder(rw).Encode(response)
			return
		}

		// Usar la libreria para validar los campos requeridos
		if validationErr := validatePoint.Struct(&point); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			_ = json.NewEncoder(rw).Encode(response)
			return
		}

		update := bson.M{
			"comment":   point.Comment,
			"address":   point.Address,
			"datetime":  point.DateTime,
			"photourl":  point.PhotoURL,
			"latitude":  point.Latitude,
			"longitude": point.Longitude,
		}

		result, err := pointCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.PointResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			_ = json.NewEncoder(rw).Encode(response)
			return
		}

		// Actualizar los campos del punto
		var updatedPoint models.Point
		if result.MatchedCount == 1 {
			err := pointCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedPoint)

			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.PointResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				_ = json.NewEncoder(rw).Encode(response)
				return
			}
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.PointResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedPoint}}
		_ = json.NewEncoder(rw).Encode(response)
		fmt.Printf("Datos del punto %s modificados con éxito\n", pointId)
	}
}

func DeleteAPoint() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		pointId := params["pointId"]
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(pointId)

		result, err := pointCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.PointResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			_ = json.NewEncoder(rw).Encode(response)
			return
		}

		if result.DeletedCount < 1 {
			rw.WriteHeader(http.StatusNotFound)
			response := responses.PointResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Point with specified ID not found!"}}
			_ = json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Point successfully deleted!"}}
		_ = json.NewEncoder(rw).Encode(response)
		fmt.Printf("Punto %s eliminado con éxito\n", pointId)
	}
}
