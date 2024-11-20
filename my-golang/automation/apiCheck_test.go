package main

import (
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {

	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal("Error creating request")
	}
}
