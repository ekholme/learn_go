package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetUser(t *testing.T) {
	s := NewServer()

	ts := httptest.NewServer(http.HandlerFunc(s.handleGetUser))
}

//RESUME HERE ~8:00 IN VIDEO
