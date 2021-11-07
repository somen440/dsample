package main

import (
	"context"
	"reflect"
	"testing"
	"time"

	"cloud.google.com/go/datastore"
)

func TestIndex(t *testing.T) {
	t.Skip()

	ctx := context.Background()
	now := time.Now()

	users := []User{
		{
			"ce6b04f2-c6bd-44fb-b9dd-25ff1ccdaacd",
			"ほげたろう",
			1011,
			now.Add(10 * time.Hour),
		},
		{
			"e4f65fe6-1d70-48cf-86cc-e33c9b34d6a5",
			"ばーじろう",
			1010,
			now.Add(10 * time.Hour),
		},
		{
			"854cb6db-9e32-48a1-a388-774cd814ea3e",
			"ふーたろう",
			1000,
			now,
		},
	}

	keys := []*datastore.Key{}
	for i := 0; i < len(users); i++ {
		keys = append(keys, datastore.IncompleteKey("Users", nil))
	}
	var err error
	keys, err = client.AllocateIDs(ctx, keys)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := client.PutMulti(ctx, keys, users); err != nil {
		t.Fatal(err)
	}
}

func TestGetIndex(t *testing.T) {
	t.Skip()

	ctx := context.Background()

	var users []User
	q := datastore.NewQuery("Users").Filter("number <=", 1010).Order("-number").Order("-created")
	if _, err := client.GetAll(ctx, q, &users); err != nil {
		t.Fatal(err)
	}

	if len(users) != 2 {
		t.Fatalf("want = 2, len(users) = %d", len(users))
	}

	expected := []int{1010, 1000}
	actual := []int{users[0].Number, users[1].Number}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("want = %+v, act = %+v", expected, actual)
	}
}
