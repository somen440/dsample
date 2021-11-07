package main

import "time"

type User struct {
	ID      string    `datastore:"id"`
	Name    string    `datastore:"name"`
	Number  int       `datastore:"number"`
	Created time.Time `datastore:"created"`
}
