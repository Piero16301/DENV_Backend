package controllers

//
//import (
//	"Deteccion_Zonas_Dengue_Backend/configs"
//	"Deteccion_Zonas_Dengue_Backend/models"
//	"Deteccion_Zonas_Dengue_Backend/responses"
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/go-playground/validator/v10"
//	"github.com/gorilla/mux"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"go.mongodb.org/mongo-driver/mongo"
//	"net/http"
//	"time"
//)
//
//var photoCollection = configs.GetCollection(configs.DB, "PropagationZone")
//var validatePhoto = validator.New()
//
//func CreatePhoto() http.HandlerFunc {
//	return func(rw http.ResponseWriter, r *http.Request) {
//		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		var photo models.Photo
//		defer cancel()
//
//		// Validar el body del requests
//		if err := json.NewDecoder(r.Body).Decode(&photo); err != nil {
//			rw.WriteHeader(http.StatusBadRequest)
//			response := responses.PhotosResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			_ = json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		// Usar la libería para validar los campos requeridos
//		if validationErr := validatePhoto.Struct(&photo); validationErr != nil {
//			rw.WriteHeader(http.StatusBadRequest)
//			response := responses.PhotosResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
//			_ = json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		newPhoto := models.Photo{
//			Id:        primitive.NewObjectID(),
//			Address:   photo.Address,
//			Comment:   photo.Comment,
//			DateTime:  photo.DateTime,
//			Latitude:  photo.Latitude,
//			Longitude: photo.Longitude,
//			PhotoURL:  photo.PhotoURL,
//		}
//
//		result, err := photoCollection.InsertOne(ctx, newPhoto)
//		if err != nil {
//			rw.WriteHeader(http.StatusInternalServerError)
//			response := responses.PhotosResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			_ = json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		rw.WriteHeader(http.StatusCreated)
//		response := responses.PhotosResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
//		_ = json.NewEncoder(rw).Encode(response)
//		fmt.Println("Nueva foto creada con éxito")
//	}
//}
//
//func GetAPhoto() http.HandlerFunc {
//	return func(rw http.ResponseWriter, r *http.Request) {
//		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		params := mux.Vars(r)
//		photoId := params["photoId"]
//		var photo models.Photo
//		defer cancel()
//
//		objId, _ := primitive.ObjectIDFromHex(photoId)
//
//		err := photoCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&photo)
//		if err != nil {
//			rw.WriteHeader(http.StatusInternalServerError)
//			response := responses.PhotosResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			_ = json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		rw.WriteHeader(http.StatusOK)
//		response := responses.PhotosResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": photo}}
//		_ = json.NewEncoder(rw).Encode(response)
//		fmt.Printf("Foto %s leída con éxito\n", photoId)
//	}
//}
//
//func EditAPhoto() http.HandlerFunc {
//	return func(rw http.ResponseWriter, r *http.Request) {
//		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		params := mux.Vars(r)
//		photoId := params["photoId"]
//		var photo models.Photo
//		defer cancel()
//
//		objId, _ := primitive.ObjectIDFromHex(photoId)
//
//		// Validar el body del request
//		if err := json.NewDecoder(r.Body).Decode(&photo); err != nil {
//			rw.WriteHeader(http.StatusBadRequest)
//			response := responses.PhotosResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			_ = json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		// Usar la librería para validar los campos requeridos
//		if validationErr := validatePhoto.Struct(&photo); validationErr != nil {
//			rw.WriteHeader(http.StatusBadRequest)
//			response := responses.PhotosResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
//			_ = json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		update := bson.M{
//			"address":   photo.Address,
//			"comment":   photo.Comment,
//			"datetime":  photo.DateTime,
//			"latitude":  photo.Latitude,
//			"longitude": photo.Longitude,
//			"photourl":  photo.PhotoURL,
//		}
//
//		result, err := photoCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
//		if err != nil {
//			rw.WriteHeader(http.StatusInternalServerError)
//			response := responses.PhotosResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			_ = json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		// Actualizar los campos de la foto
//		var updatedPhoto models.Photo
//		if result.MatchedCount == 1 {
//			err := photoCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedPhoto)
//			if err != nil {
//				rw.WriteHeader(http.StatusInternalServerError)
//				response := responses.PhotosResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//				_ = json.NewEncoder(rw).Encode(response)
//				return
//			}
//		}
//
//		rw.WriteHeader(http.StatusOK)
//		response := responses.PhotosResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedPhoto}}
//		_ = json.NewEncoder(rw).Encode(response)
//		fmt.Printf("Datos de la foto %s modificados con éxito\n", photoId)
//	}
//}
//
//func DeleteAPhoto() http.HandlerFunc {
//	return func(rw http.ResponseWriter, r *http.Request) {
//		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		params := mux.Vars(r)
//		photoId := params["photoId"]
//		defer cancel()
//
//		objId, _ := primitive.ObjectIDFromHex(photoId)
//
//		result, err := photoCollection.DeleteOne(ctx, bson.M{"id": objId})
//		if err != nil {
//			rw.WriteHeader(http.StatusInternalServerError)
//			response := responses.PhotosResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			_ = json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		if result.DeletedCount < 1 {
//			rw.WriteHeader(http.StatusNotFound)
//			response := responses.PhotosResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Photo with specified ID not found!"}}
//			_ = json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		rw.WriteHeader(http.StatusOK)
//		response := responses.PhotosResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Photo successfully deleted!"}}
//		_ = json.NewEncoder(rw).Encode(response)
//		fmt.Printf("Foto %s eliminado con éxito\n", photoId)
//	}
//}
//
//func GetAllPhotos() http.HandlerFunc {
//	return func(rw http.ResponseWriter, r *http.Request) {
//		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		var photos []models.Photo
//		defer cancel()
//
//		results, err := photoCollection.Find(ctx, bson.M{})
//
//		if err != nil {
//			rw.WriteHeader(http.StatusInternalServerError)
//			response := responses.PhotosResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			_ = json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		// Lectura de manera óptima de la BD
//		defer func(results *mongo.Cursor, ctx context.Context) {
//			_ = results.Close(ctx)
//		}(results, ctx)
//
//		for results.Next(ctx) {
//			var singlePhoto models.Photo
//			if err = results.Decode(&singlePhoto); err != nil {
//				rw.WriteHeader(http.StatusInternalServerError)
//				response := responses.PhotosResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//				_ = json.NewEncoder(rw).Encode(response)
//			}
//			photos = append(photos, singlePhoto)
//		}
//
//		rw.WriteHeader(http.StatusOK)
//		response := responses.PhotosResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": photos}}
//		_ = json.NewEncoder(rw).Encode(response)
//		fmt.Println("Fotos leídas con éxito")
//	}
//}
//
//func DeleteAllPhotos() http.HandlerFunc {
//	return func(rw http.ResponseWriter, r *http.Request) {
//		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		defer cancel()
//
//		_, err := photoCollection.DeleteMany(ctx, bson.M{})
//		if err != nil {
//			rw.WriteHeader(http.StatusInternalServerError)
//			response := responses.PhotosResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
//			_ = json.NewEncoder(rw).Encode(response)
//			return
//		}
//
//		rw.WriteHeader(http.StatusOK)
//		response := responses.PhotosResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Photos successfully deleted!"}}
//		_ = json.NewEncoder(rw).Encode(response)
//		fmt.Println("Fotos eliminadas con éxito")
//	}
//}
