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

func dump(r *http.Request) {
	output, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
		return
	}
	fmt.Println(string(output))
}

func CreateFoodstuffEndpoint(response http.ResponseWriter, request *http.Request) {
	dump(request)
	newFoodstuff := Foodstuff{"", "Brussel Sprouts", "roasted yumminess", []string{"leafy veggie balls"}}
	dbContext, err := GetDbContext("")
	fmt.Fprintf(os.Stdout, "Inserting Foodstuff %v", newFoodstuff)
	result, e2 := dbContext.GetDbCollection().InsertOne(dbContext.Context, newFoodstuff) // INSERT
	if e2 != nil {
		fmt.Println(fmt.Sprintf("Error from InsertOne(): %v", e2))
		os.Exit(1);
	} else {
		fmt.Println(fmt.Sprintf("Successfully inserted foodstuff with ID %v, Result: %v", newFoodstuff.ID, result))
	}
}

func GetFoodstuffEndpoint(response http.ResponseWriter, request *http.Request) {
	dump(request)
	fmt.Fprintf(os.Stdout, "Getting Foodstuff %s", request)
}

func main() {
	fmt.Println("Starting the food-api...")
	router := mux.NewRouter()
	router.HandleFunc("/foodstuff", CreateFoodstuffEndpoint).Methods("POST")
	router.HandleFunc("/foodstuff/{id}", GetFoodstuffEndpoint).Methods("GET")
	log.fatal(http.ListenAndServe(":12345", router))
}
