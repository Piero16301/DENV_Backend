package controllers

import (
	"Deteccion_Zonas_Dengue_Backend/configs"
	"Deteccion_Zonas_Dengue_Backend/models"
	"Deteccion_Zonas_Dengue_Backend/responses"
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
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
	}
}
