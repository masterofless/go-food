package main

import (
	"fmt"
	"testing"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func TestDatabaseConnection(t *testing.T) {
	var client *mongo.Client
	url := "mongodb://food-db-server:27017"
	db := "food-db"
	collectionName := "foodstuffs"
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	clientOptions := options.Client().ApplyURI(url)
	client, _ = mongo.Connect(ctx, clientOptions)
	collection := client.Database(db).Collection(collectionName)

	foodstuff := NewFoodstuff()
	result, err := collection.InsertOne(ctx, foodstuff)
	if err != nil {
		t.Error(fmt.Sprintf("Error from InsertOne(): %v", err))
	}
	fmt.Println(fmt.Sprintf("Inserted foodstuff Result: %v, %v", result, foodstuff))

	id, _ := primitive.ObjectIDFromHex("5d2399ef96fb765873a24bae")
	fmt.Println(fmt.Sprintf("id from Hex: %v", id))
	var f2 Foodstuff
	e2 := collection.FindOne(ctx, Foodstuff{ID: id}).Decode(&f2)
	if e2 != nil {
		t.Error(fmt.Sprintf("Unable to FindOne() foodstuff %v", e2))
	}
	fmt.Println(fmt.Sprintf("got result from findOne: %v", f2))
}

func TestFoodstuff(t *testing.T) {
	want := NewFoodstuff()
	foodstuff := NewFoodstuff()
	if foodstuff.CanonicalName != want.CanonicalName {
		t.Errorf("Not the food I was looking for")
	}
}
