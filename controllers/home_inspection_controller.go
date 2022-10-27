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
			Larvicide:         homeInspection.Larvicide,
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
		// return id of created home inspection as an id field
		response := responses.HomeInspectionResponse{
			Status:  http.StatusCreated,
			Message: "Inspección de vivienda creada con éxito",
			Data:    map[string]interface{}{"id": newHomeInspection.Id},
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
			"larvicide":         homeInspection.Larvicide,
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
			Data:    map[string]interface{}{"id": updatedHomeInspection.Id},
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
		params := mux.Vars(request)
		skip, _ := strconv.ParseInt(params["skip"], 0, 64)
		var homeInspections []models.HomeInspection
		defer cancel()

		results, err := homeInspectionCollection.Find(ctx, bson.M{}, options.Find().SetSkip(skip))

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
		params := mux.Vars(request)
		skip, _ := strconv.ParseInt(params["skip"], 0, 64)
		var homeInspectionsSummarized []models.HomeInspectionSummarized
		defer cancel()

		results, err := homeInspectionCollection.Find(ctx, bson.M{}, &options.FindOptions{Projection: bson.M{
			"id":        1,
			"latitude":  1,
			"longitude": 1,
			"datetime":  1,
			"photourl":  1,
		}}, options.Find().SetSkip(skip))

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
			var singleHomeInspectionSummarized models.HomeInspectionSummarized
			if err = results.Decode(&singleHomeInspectionSummarized); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.HomeInspectionResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al obtener las inspecciones de vivienda resumidas",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
			}
			homeInspectionsSummarized = append(homeInspectionsSummarized, singleHomeInspectionSummarized)
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusOK,
			Message: "Inspecciones de vivienda resumidas obtenidas con éxito",
			Data:    homeInspectionsSummarized,
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

func GetHomeInspectionClusters() http.HandlerFunc {
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

		// Parse date from string to time.Time 2022-09-08T12:10:26.000+00:00
		startDateParsed, err := time.Parse(time.RFC3339, startDate)
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

		endDateParsed, err := time.Parse(time.RFC3339, endDate)
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

		results, err := homeInspectionCollection.Find(ctx, bson.M{
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
			response := responses.HomeInspectionResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ocurrió un error al leer las inspecciones de vivienda",
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
			var singleHomeInspection models.HomeInspectionSummarized
			if err = results.Decode(&singleHomeInspection); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				response := responses.HomeInspectionResponse{
					Status:  http.StatusInternalServerError,
					Message: "Ocurrió un error al leer las inspecciones de vivienda",
					Data:    err.Error(),
				}
				_ = json.NewEncoder(writer).Encode(response)
			}
			dataPoints = append(
				dataPoints,
				&dbs.NamedPoint{
					Name: singleHomeInspection.Id.Hex(),
					Point: []float64{
						float64(singleHomeInspection.Latitude),
						float64(singleHomeInspection.Longitude),
					},
					Date: singleHomeInspection.Datetime,
				})
		}

		var clusterResult = clusterer.Cluster(dataPoints)
		var clusters []models.Cluster

		for i, cluster := range clusterResult {
			var clusterPoints []models.ClusterPoint
			for _, point := range cluster {
				clusterPoints = append(clusterPoints, models.ClusterPoint{
					Id:        point.GetName(),
					Latitude:  float32(point.GetPoint()[0]),
					Longitude: float32(point.GetPoint()[1]),
					Datetime:  point.GetDate(),
				})
			}
			clusters = append(clusters, models.Cluster{
				Id:     i + 1,
				Points: clusterPoints,
			})
		}

		writer.WriteHeader(http.StatusOK)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusOK,
			Message: "Clusters de inspecciones de vivienda obtenidos con éxito",
			Data:    clusters,
		}
		_ = json.NewEncoder(writer).Encode(response)
		fmt.Println("Clusters de inspecciones de vivienda obtenidos con éxito")
	}
}
