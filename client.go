package main

import (
	"context"

	"cloud.google.com/go/datastore"
)

const projectID = "my-datastore-project"

var client *datastore.Client

func init() {
	var err error
	client, err = datastore.NewClient(context.TODO(), projectID)
	if err != nil {
		panic(err)
	}
}
