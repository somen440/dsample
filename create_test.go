package main

import (
	"context"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()

	user := &User{
		ID:   uuid.New().String(),
		Name: "テストマン",
	}

	projectID := "my-datastore-project"

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		t.Fatal(err)
	}

	k := datastore.NameKey("Users", uuid.New().String(), nil)
	if _, err := client.Put(ctx, k, user); err != nil {
		t.Fatal(err)
	}
}
