package controllers

import (
	"Deteccion_Zonas_Dengue_Backend/configs"
	"Deteccion_Zonas_Dengue_Backend/models"
	"Deteccion_Zonas_Dengue_Backend/responses"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

var photoCollection = configs.GetCollection(configs.DB, "photos")
var validatePhoto = validator.New()

func CreatePhoto() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var photo models.Photo
		defer cancel()

		// Validar el body del requests
		if err := json.NewDecoder(r.Body).Decode(&photo); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.PhotosResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			_ = json.NewEncoder(rw).Encode(response)
			return
		}

		// Usar la libería para validar los campos requeridos
		if validationErr := validatePhoto.Struct(&photo); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.PhotosResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			_ = json.NewEncoder(rw).Encode(response)
			return
		}

		newPhoto := models.Photo{
			Id:        primitive.NewObjectID(),
			Address:   photo.Address,
			Comment:   photo.Comment,
			DateTime:  photo.DateTime,
			Latitude:  photo.Latitude,
			Longitude: photo.Longitude,
			PhotoURL:  photo.PhotoURL,
		}

		result, err := photoCollection.InsertOne(ctx, newPhoto)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.PhotosResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			_ = json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.PhotosResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		_ = json.NewEncoder(rw).Encode(response)
		fmt.Println("Nueva foto creada con éxito")
	}
}
