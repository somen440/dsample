package main

import (
	"context"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()

	t.Skip()

	user := &User{
		ID:   uuid.New().String(),
		Name: "テストマン",
	}

	k := datastore.NameKey("Users", uuid.New().String(), nil)
	if _, err := client.Put(ctx, k, user); err != nil {
		t.Fatal(err)
	}
}
