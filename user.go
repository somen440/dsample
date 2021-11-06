package main

type User struct {
	ID   string `datastore:"id"`
	Name string `datastore:"name"`
}
