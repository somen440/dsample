package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()

	client.Close()

	os.Exit(code)
}
