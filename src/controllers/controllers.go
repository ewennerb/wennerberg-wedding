package controllers

import (
	"context"
	"encoding/json"
	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"wennerbeg-wedding/db"
	"wennerbeg-wedding/models"
)

var client = db.ConnectDB()

var GetHouseholdByGuest = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	// Parse Params
	params := mux.Vars(request)
	name := params["guest_name"]
	logger := log.New(os.Stderr)
	logger.Infof("GET /household/%s", name)

	// Query DB for the household we're looking for
	var household models.Household
	collection := client.Database("test").Collection("household")
	err := collection.FindOne(context.TODO(), bson.M{"names": name}).Decode(&household)

	//Return a 404 if the guest isn't found
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(err.Error()))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(household)
})

var UpdateHousehold = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	logger := log.New(os.Stderr)
	logger.Infof("POST /household/%s", id)

	var household models.Household
	json.NewDecoder(request.Body).Decode(&household)

	collection := client.Database("test").Collection("household")
	_, err := collection.ReplaceOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}, household)

	//Todo: Assert result.ModifiedCount == 1?
	// Todo: Better error handling
	// Todo: Add 400 error if dinner choice is invalid
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": Could not save guest preferences"`))
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	//err = json.NewEncoder(response).Encode(household)
	//Todo: Return 201 response
})

var GetGuestInfo = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	collection := client.Database("test").Collection("household")
	pipeline := mongo.Pipeline{
		{{"$unwind", bson.D{{"path", "$guests"}}}},
		{
			{"$group", bson.D{
				{"_id", nil},
				{"invited", bson.D{{"$sum", 1}}},
				{"attending", bson.D{{"$sum", bson.D{{"$cond", bson.A{bson.D{{"$eq", bson.A{"$guests.attending", true}}}, 1, 0}}}}}},
				{"steak", bson.D{{"$sum", bson.D{{"$cond", bson.A{bson.D{{"$eq", bson.A{"$guests.dinner_choice", "steak"}}}, 1, 0}}}}}},
				{"fish", bson.D{{"$sum", bson.D{{"$cond", bson.A{bson.D{{"$eq", bson.A{"$guests.dinner_choice", "fish"}}}, 1, 0}}}}}},
				{"lasagna", bson.D{{"$sum", bson.D{{"$cond", bson.A{bson.D{{"$eq", bson.A{"$guests.dinner_choice", "lasagna"}}}, 1, 0}}}}}},
				{"attending_guests", bson.D{{"$push", bson.D{{"$cond", bson.A{
					bson.D{{"$eq", bson.A{"$guests.attending", true}}},
					bson.D{
						{"name", "$guests.name"},
						{"dinner", "$guests.dinner_choice"},
					},
					nil,
				}}}}}},
			}},
		},
		{
			{"$project", bson.D{
				{"_id", 0},
				{"invited", 1},
				{"attending", 1},
				{"steak", 1},
				{"fish", 1},
				{"lasagna", 1},
				{"attending_guests", bson.D{{"$filter", bson.D{
					{"input", "$attending_guests"},
					{"as", "guest"},
					{"cond", bson.D{{"$ne", bson.A{"$$guest", nil}}}},
				}}}},
			}},
		},
	}
	// Retrieve aggregation of dinners, guests, and my guest's selections
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		http.Error(response, "Failed to retrieve aggregate data from DB", http.StatusInternalServerError)
		return
	}

	// Decode the response from the DB to BSON - raise a 500 if this fails
	var result bson.M
	if cursor.Next(context.Background()) {
		if err := cursor.Decode(&result); err != nil {
			http.Error(response, "Failed to decode result", http.StatusInternalServerError)
			return
		}
	}

	// Encode JSON Response - raise a 500 if this fails for some reason
	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(result); err != nil {
		http.Error(response, "Failed to encode JSON Response", http.StatusInternalServerError)
		return
	}
})
