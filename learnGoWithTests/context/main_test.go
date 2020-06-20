package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubStore struct {
	response string
}

func (s *StubStore) fetch() string {
	return s.response
}

func TestHandler(t *testing.T) {
	data := "olá, mundo"
	svr := server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`resultado "%s", esperado "%s"`, response.Body.String(), data)
	}
}
