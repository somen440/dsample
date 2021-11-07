package main

import (
	"context"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
)

func TestNamespace(t *testing.T) {
	t.Skip()

	ctx := context.Background()

	t.Run("empty namespace", func(t *testing.T) {
		k := datastore.IncompleteKey("Users", nil)
		if k.Namespace != "" {
			t.Fatal("namespace is not empty", k.Namespace)
		}

		user := &User{
			ID:   uuid.New().String(),
			Name: "namespace empty テストマン",
		}

		if _, err := client.Put(ctx, k, user); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("specified namespace", func(t *testing.T) {
		k := datastore.IncompleteKey("Users", nil)
		k.Namespace = "test"

		user := &User{
			ID:   uuid.New().String(),
			Name: "namespace test テストマン",
		}

		if _, err := client.Put(ctx, k, user); err != nil {
			t.Fatal(err)
		}
	})
}
