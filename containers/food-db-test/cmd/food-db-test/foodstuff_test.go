package main

import (
	"fmt"
	"testing"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
	"os"
)

func TestPing(t *testing.T) {
	url := fmt.Sprintf("mongodb://%s:%s@food-db-server:27017/?authSource=admin", os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"))
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	collection := client.Database("testing").Collection("numbers")
	if collection == nil {
		t.Error("that did not go well")
	} else {
		fmt.Println("Test ping went well")
	}
	err = client.Connect(ctx)
	if err != nil {
		t.Error("client.Connect did not go well")
	} else {
		fmt.Println("client.Connect went well")
	}
}

func testInsertFoodstuff(t *testing.T, dbContext *DbContext) string {
	t.Helper()
	newFoodstuff := Foodstuff{"", "Bread", "Earth", []string{"Mana"}}
	result, err := dbContext.GetDbCollection().InsertOne(dbContext.Context, newFoodstuff)
	if err != nil {
		t.Error(fmt.Sprintf("Error from InsertOne(): %v", err))
	} else {
		fmt.Println(fmt.Sprintf("Successfully inserted foodstuff with ID %v, Result: %v", newFoodstuff.ID, result))
	}
	return newFoodstuff.ID
}

func testReadFoodstuff(t *testing.T, dbContext *DbContext, id string) {
	t.Helper()
	var doneReadFoodstuff Foodstuff

	fmt.Println(fmt.Sprintf("FindOne foodstuff with id %s", id))
	err := dbContext.GetDbCollection().FindOne(dbContext.Context, Foodstuff{ID: id}).Decode(&doneReadFoodstuff)
	if err != nil {
		t.Error(fmt.Sprintf("Unable to FindOne() foodstuff %v", err))
	} else if len(doneReadFoodstuff.Description) == 0 {
		fmt.Println(fmt.Sprintf("No Description"))
	} else {
		fmt.Println(fmt.Sprintf("got Description from findOne: %s", doneReadFoodstuff.Description))
	}
}

func TestDatabaseUsage(t *testing.T) {
	dbContext, _ := GetDbContext("")
	id := testInsertFoodstuff(t, &dbContext)
	testReadFoodstuff(t, &dbContext, id)
}
