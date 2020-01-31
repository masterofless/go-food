package main

import (
	"fmt"
	"os"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Foodstuff struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CanonicalName string             `json:"canonicalname,omitempty" bson:"canonicalname,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
	AKA           []string           `json:"aka,omitempty" bson:"aka,omitempty"`
}

func NewFoodstuff() Foodstuff {
	f := Foodstuff{primitive.NewObjectID(), "Brussel Sprouts", "roasted yumminess", []string{"leafy veggie balls"}}
	return f
}

func (f *Foodstuff) ToString() string {
	return "My Foodname is " + f.CanonicalName
}

func main() {
	f := NewFoodstuff()
	fmt.Fprintf(os.Stdout, "Starting Foodstuff %v", f)
}
