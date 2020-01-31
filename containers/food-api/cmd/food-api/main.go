package main

import (
	"context"
	// "encoding/json"
	"fmt"
	// "log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Foodstuff struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CanonicalName string             `json:"canonicalname,omitempty" bson:"canonicalname,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
	AKA           string             `json:"aka,omitempty" bson:"aka,omitempty"`
}

func CreateFoodstuffEndpoint(response http.ResponseWriter, request *http.Request) {}
func GetRecipesEndpoint(response http.ResponseWriter, request *http.Request) { }
func GetFoodstuffEndpoint(response http.ResponseWriter, request *http.Request) { }

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://food-db-server:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/foodstuff", CreateFoodstuffEndpoint).Methods("POST")
	router.HandleFunc("/recipes", GetRecipesEndpoint).Methods("GET")
	router.HandleFunc("/foodstuff/{id}", GetFoodstuffEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}
