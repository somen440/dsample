package main

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/google/go-cmp/cmp"
)

func TestGet(t *testing.T) {
	ctx := context.Background()

	t.Skip()

	var users []User
	if _, err := client.GetAll(ctx, datastore.NewQuery("Users"), &users); err != nil {
		t.Fatal(err)
	}

	if len(users) != 3 {
		t.Fatalf("want=3, act=%d", len(users))
	}
	expected := []User{
		{
			"ce6b04f2-c6bd-44fb-b9dd-25ff1ccdaacd",
			"ほげたろう",
			1,
			time.Now(),
		},
		{
			"e4f65fe6-1d70-48cf-86cc-e33c9b34d6a5",
			"ばーじろう",
			1,
			time.Now(),
		},
		{
			"854cb6db-9e32-48a1-a388-774cd814ea3e",
			"ふーたろう",
			1,
			time.Now(),
		},
	}
	if d := cmp.Diff(expected, users); d != "" {
		t.Errorf("got = %v, want %v, diff %s", users, expected, d)
	}
}
