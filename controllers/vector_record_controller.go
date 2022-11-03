package controllers

import (
	"DENV_Backend/configs"
	dbs "DENV_Backend/dbscan"
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
	"strconv"
	"time"
)

var vectorRecordCollection = configs.GetCollection(configs.DB, "VectorRecord")
var validateVectorRecord = validator.New()

func CreateVectorRecord() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var vectorRecord models.VectorRecord
		defer cancel()

		// Validar que el body está en formato JSON
		if err := json.NewDecoder(request.Body).Decode(&vectorRecord); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.VectorRecordResponse{
				Status:  http.StatusBadRequest,
				Message: "El cuerpo de la solicitud no está en formato JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Se valida que se envíen todos los campos requeridos
		if validationErr := validateVectorRecord.Struct(&vectorRecord); validationErr != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.VectorRecordResponse{
				Status:  http.StatusBadRequest,
				Message: "No se han enviado todos los campos requeridos",
				Data:    validationErr.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Se crea instancia de registro de vector
		newVectorRecord := models.VectorRecord{
			Id:        primitive.NewObjectID(),
			Address:   vectorRecord.Address,
			Comment:   vectorRecord.Comment,
			Datetime:  vectorRecord.Datetime,
			Latitude:  vectorRecord.Latitude,
			Longitude: vectorRecord.Longitude,
			PhotoURL:  vectorRecord.PhotoURL,
		}

		// Se inserta el registro de vector en MongoDB
		_, err := vectorRecordCollection.InsertOne(ctx, newVectorRecord)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.VectorRecordResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al crear el registro de vector",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusCreated)
		response := responses.VectorRecordResponse{
			Status:  http.StatusCreated,
			Message: "Registro de vector creado con éxito",
			Data:    map[string]interface{}{"id": newVectorRecord.Id},
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Registro de vector %s creado con éxito\n", newVectorRecord.Id.Hex())
	}
}

func GetVectorRecord() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(request)
		vectorRecordId := params["vectorRecordId"]
		var vectorRecord models.VectorRecord
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(vectorRecordId)

		err := vectorRecordCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&vectorRecord)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.VectorRecordResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al obtener el registro de vector",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.VectorRecordResponse{
			Status:  http.StatusOK,
			Message: "Registro de vector obtenido con éxito",
			Data:    vectorRecord,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Registro de vector %s obtenido con éxito\n", vectorRecordId)
	}
}

func EditVectorRecord() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(request)
		vectorRecordId := params["vectorRecordId"]
		var vectorRecord models.VectorRecord
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(vectorRecordId)

		// Validar que el body está en formato JSON
		if err := json.NewDecoder(request.Body).Decode(&vectorRecord); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.VectorRecordResponse{
				Status:  http.StatusBadRequest,
				Message: "El cuerpo de la solicitud no está en formato JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Se valida que se envíen todos los campos requeridos
		if validationErr := validateVectorRecord.Struct(&vectorRecord); validationErr != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.VectorRecordResponse{
				Status:  http.StatusBadRequest,
				Message: "No se han enviado todos los campos requeridos",
				Data:    validationErr.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		bsonVectorRecord := bson.M{
			"address":   vectorRecord.Address,
			"comment":   vectorRecord.Comment,
			"datetime":  vectorRecord.Datetime,
			"latitude":  vectorRecord.Latitude,
			"longitude": vectorRecord.Longitude,
			"photourl":  vectorRecord.PhotoURL,
		}

		result, err := vectorRecordCollection.UpdateOne(ctx, bson.M{"id": objectId}, bson.M{"$set": bsonVectorRecord})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.VectorRecordResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al actualizar el registro de vector",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Actualizar los campos del registro de vector
		var updatedVectorRecord models.VectorRecord
		if result.MatchedCount == 1 {
			err := vectorRecordCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&updatedVectorRecord)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.VectorRecordResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al actualizar el registro de vector",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
				return
			}
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.VectorRecordResponse{
			Status:  http.StatusOK,
			Message: "Registro de vector actualizado con éxito",
			Data:    map[string]interface{}{"id": updatedVectorRecord.Id},
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Registro de vector %s actualizado con éxito\n", vectorRecordId)
	}
}

func DeleteVectorRecord() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(request)
		vectorRecordId := params["vectorRecordId"]
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(vectorRecordId)

		result, err := vectorRecordCollection.DeleteOne(ctx, bson.M{"id": objectId})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.VectorRecordResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al eliminar el registro de vector",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		if result.DeletedCount < 1 {
			writer.WriteHeader(http.StatusNotFound)
			response := responses.VectorRecordResponse{
				Status:  http.StatusNotFound,
				Message: "No se encontró el registro de vector con el ID especificado",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.VectorRecordResponse{
			Status:  http.StatusOK,
			Message: "Registro de vector eliminado con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Printf("Registro de vector %s eliminado con éxito\n", vectorRecordId)
	}
}

func GetAllVectorRecordsDetailed() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		params := mux.Vars(request)
		skip, _ := strconv.ParseInt(params["skip"], 0, 64)

		startDate := request.URL.Query().Get("startDate")
		endDate := request.URL.Query().Get("endDate")

		var vectorRecords []models.VectorRecord

		defer cancel()

		if startDate == "" || endDate == "" {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "Fechas de inicio o fin no están correctamente especificadas",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Parse date from string to time.Time 2022-10-01T22:49:16.072255
		startDateParsed, err := time.Parse("2006-01-02T15:04:05.000000", startDate)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "Fecha de inicio no está correctamente especificada",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		endDateParsed, err := time.Parse("2006-01-02T15:04:05.000000", endDate)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "Fecha de fin no está correctamente especificada",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		results, err := vectorRecordCollection.Find(ctx, bson.M{
			"datetime": bson.M{
				"$gte": startDateParsed,
				"$lte": endDateParsed,
			},
		}, options.Find().SetSkip(skip))

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.VectorRecordResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al obtener los registros de vector detallados",
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
			var singleVectorRecord models.VectorRecord
			if err = results.Decode(&singleVectorRecord); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.VectorRecordResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al obtener los registros de vector detallados",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
			}
			vectorRecords = append(vectorRecords, singleVectorRecord)
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.VectorRecordResponse{
			Status:  http.StatusOK,
			Message: "Registros de vector detallados obtenidos con éxito",
			Data:    vectorRecords,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Println("Registros de vector detallados obtenidos con éxito")
	}
}

func GetAllVectorRecordsSummarized() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		params := mux.Vars(request)
		skip, _ := strconv.ParseInt(params["skip"], 0, 64)

		startDate := request.URL.Query().Get("startDate")
		endDate := request.URL.Query().Get("endDate")

		var vectorRecordsSummarized []models.VectorRecordSummarized

		defer cancel()

		if startDate == "" || endDate == "" {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "Fechas de inicio o fin no están correctamente especificadas",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Parse date from string to time.Time 2022-10-01T22:49:16.072255
		startDateParsed, err := time.Parse("2006-01-02T15:04:05.000000", startDate)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "Fecha de inicio no está correctamente especificada",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		endDateParsed, err := time.Parse("2006-01-02T15:04:05.000000", endDate)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "Fecha de fin no está correctamente especificada",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		results, err := vectorRecordCollection.Find(ctx, bson.M{
			"datetime": bson.M{
				"$gte": startDateParsed,
				"$lte": endDateParsed,
			},
		}, &options.FindOptions{Projection: bson.M{
			"id":        1,
			"latitude":  1,
			"longitude": 1,
			"datetime":  1,
			"photourl":  1,
		}}, options.Find().SetSkip(skip))

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.VectorRecordResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al obtener los registros de vector resumidos",
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
			var singleVectorRecordSummarized models.VectorRecordSummarized
			if err = results.Decode(&singleVectorRecordSummarized); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.VectorRecordResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al obtener los registros de vector resumidos",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
			}
			vectorRecordsSummarized = append(vectorRecordsSummarized, singleVectorRecordSummarized)
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.VectorRecordResponse{
			Status:  http.StatusOK,
			Message: "Registros de vector resumidos obtenidos con éxito",
			Data:    vectorRecordsSummarized,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Println("Registros de vector resumidos obtenidos con éxito")
	}
}

func DeleteAllVectorRecords() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err := vectorRecordCollection.DeleteMany(ctx, bson.M{})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.VectorRecordResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al eliminar todos los registros de vector",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.VectorRecordResponse{
			Status:  http.StatusOK,
			Message: "Todos los registros de vector eliminados con éxito",
			Data:    nil,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Println("Registros de vector eliminados con éxito")
	}
}

func GetVectorRecordClusters() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		// Increase the timeout to 30 seconds
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

		params := mux.Vars(request)
		eps, _ := strconv.ParseFloat(params["eps"], 64)
		minPoints, _ := strconv.Atoi(params["minPoints"])

		startDate := request.URL.Query().Get("startDate")
		endDate := request.URL.Query().Get("endDate")

		defer cancel()

		if startDate == "" || endDate == "" {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "Fechas de inicio o fin no están correctamente especificadas",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		if eps <= 0 || minPoints <= 0 {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "Los parámetros EPS y Mínimo de Puntos deben ser diferentes a cero",
				Data:    nil,
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Parse date from string to time.Time 2022-10-01T22:49:16.072255
		startDateParsed, err := time.Parse("2006-01-02T15:04:05.000000", startDate)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "Fecha de inicio no está correctamente especificada",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		endDateParsed, err := time.Parse("2006-01-02T15:04:05.000000", endDate)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "Fecha de fin no está correctamente especificada",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		results, err := vectorRecordCollection.Find(ctx, bson.M{
			"datetime": bson.M{
				"$gte": startDateParsed,
				"$lte": endDateParsed,
			},
		}, &options.FindOptions{Projection: bson.M{
			"id":        1,
			"latitude":  1,
			"longitude": 1,
			"datetime":  1,
			"photourl":  1,
		}})

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.VectorRecordResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al leer los registros de vector",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Lectura de resultados de la base de datos
		defer func(results *mongo.Cursor, ctx context.Context) {
			_ = results.Close(ctx)
		}(results, ctx)

		// Se crea el cluster con los parámetros de entrada
		var clusterer = dbs.NewDBSCANClusterer(eps, minPoints)
		var dataPoints []dbs.ClusterablePoint

		for results.Next(ctx) {
			var singleVectorRecord models.VectorRecordSummarized
			if err = results.Decode(&singleVectorRecord); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.VectorRecordResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al leer los registros de vector",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
			}
			dataPoints = append(
				dataPoints,
				&dbs.NamedPoint{
					Name: singleVectorRecord.Id.Hex(),
					Point: []float64{
						float64(singleVectorRecord.Latitude),
						float64(singleVectorRecord.Longitude),
					},
				})
		}

		var clusterResult = clusterer.Cluster(dataPoints)
		var clusters []models.Cluster

		for i, cluster := range clusterResult {
			var clusterPoints []models.ClusterPoint
			for _, point := range cluster {
				clusterPoints = append(clusterPoints, models.ClusterPoint{
					Id:        point.(*dbs.NamedPoint).Name,
					Latitude:  float32(point.GetPoint()[0]),
					Longitude: float32(point.GetPoint()[1]),
				})
			}
			clusters = append(clusters, models.Cluster{
				Id:     i + 1,
				Points: clusterPoints,
			})
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.VectorRecordResponse{
			Status:  http.StatusOK,
			Message: "Clusters de registros de vectores obtenidos con éxito",
			Data:    clusters,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Println("Clusters de registros de vector obtenidos con éxito")
	}
}
