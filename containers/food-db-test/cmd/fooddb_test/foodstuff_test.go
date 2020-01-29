package main

import (
	"fmt"
	"testing"
    gocb "github.com/couchbase/gocb/v2"
)

func TESTCBAPI() {
    opts := gocb.ClusterOptions{
        Authenticator: gocb.PasswordAuthenticator{
            "Administrator",
            "password",
        },
    }
    cluster, err := gocb.Connect("couchbase", opts)
    if err != nil {
        panic(err)
    }
    // get a bucket reference
    bucket := cluster.Bucket("travel-sample", nil)
    // get a collection reference
    collection := bucket.DefaultCollection()
    // for a named collection and scope
    // collection := bucket.Scope("my-scope").Collection("my-collection", &gocb.CollectionOptions{})

    // Upsert Document
    upsertResult, err := collection.Upsert("my-document", map[string]string{"name": "mike"}, &gocb.UpsertOptions{})
    if err != nil {
        panic(err)
    }
    fmt.Println(upsertResult)

    // Get Document
    getResult, err := collection.Get("my-document", &gocb.GetOptions{})
    if err != nil {
        panic(err)
    }
    fmt.Println(getResult)
}

func TestFoodstuff(t *testing.T) {
    want := Foodstuff{"Brussel Sprouts"}
    foodstuff := newFoodstuff()
    if foodstuff != want {
        t.Errorf("Not the food I was looking for")
    }
}
