package main

import (
	"fmt"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"time"
)

type Foodstuff struct {
	ID            string   `json:"_id,omitempty" bson:"_id,omitempty"`
	CanonicalName string   `json:"canonicalname,omitempty" bson:"canonicalname,omitempty"`
	Description   string   `json:"description,omitempty" bson:"description,omitempty"`
	AKA           []string `json:"aka,omitempty" bson:"aka,omitempty"`
}

func (f *Foodstuff) ToString() string {
	return "My Foodname is " + f.CanonicalName
}

type DbContext struct {
	Url string
	Client *mongo.Client
	Context context.Context
	// Collection *mongo.Collection
}

func (dbContext *DbContext) GetDbCollection() *mongo.Collection {
	dbName := "food-db"
	collectionName := "foodstuffs"
	return dbContext.Client.Database(dbName).Collection(collectionName)
}
func GetDbContext(url string) (c DbContext, err error) {
	if len(url) == 0 {
		url = fmt.Sprintf("mongodb://%s:%s@food-db-server:27017/", os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"))
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	// defer cancel()
	fmt.Println(fmt.Sprintf("Connecting to %v", url))
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error connecting to %v", url))
	}
	dbContext := DbContext{Url: url, Client: client, Context: ctx}
	return dbContext, err
}

func main() {
	newFoodstuff := Foodstuff{"", "Brussel Sprouts", "roasted yumminess", []string{"leafy veggie balls"}}
	dbContext, err := GetDbContext("")
	fmt.Fprintf(os.Stdout, "Inserting Foodstuff %v", newFoodstuff)
	result, e2 := dbContext.GetDbCollection().InsertOne(dbContext.Context, newFoodstuff)
	if e2 != nil {
		fmt.Println(fmt.Sprintf("Error from InsertOne(): %v", e2))
		os.Exit(1);
	} else {
		fmt.Println(fmt.Sprintf("Successfully inserted foodstuff with ID %v, Result: %v", newFoodstuff.ID, result))
	}
	fmt.Println(fmt.Sprintf("Reading based on foodstuff %v", newFoodstuff))
	var didReadFoodstuff Foodstuff
	err = dbContext.GetDbCollection().FindOne(dbContext.Context, Foodstuff{ID: newFoodstuff.ID}).Decode(&didReadFoodstuff)
	if err != nil {
		fmt.Println(fmt.Sprintf("Unable to FindOne() foodstuff %v", err))
		os.Exit(2);
	} else {
		fmt.Println(fmt.Sprintf("got result from findOne: %v", didReadFoodstuff))
	}
}
