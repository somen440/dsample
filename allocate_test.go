package main

import (
	"context"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
)

func TestAllocate(t *testing.T) {
	ctx := context.Background()

	k := datastore.IncompleteKey("Users", nil)
	k.Namespace = "test_allocate"

	keys, err := client.AllocateIDs(ctx, []*datastore.Key{k})
	if err != nil {
		t.Fatal(err)
	}

	user := &User{
		ID:   uuid.New().String(),
		Name: "allocate テストマン",
	}
	if _, err := client.Put(ctx, keys[0], user); err != nil {
		t.Fatal(err)
	}
}
