package main

import (
	"go-workshops/012-testing/04-testing-http-server/handler"

	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimpleHandler(t *testing.T) {
	// Used in E2E tests, configurable
	testServer := httptest.NewServer(http.HandlerFunc(handler.Simple))
	defer testServer.Close()

	response, err := http.Get(testServer.URL)
	if err != nil {
		log.Fatal(err)
	}

	greeting, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if string(greeting) != "Hello World" {
		t.Errorf(
			"Invalid greeting!\n I've got \t%s\n but want: \t%s",
			string(greeting),
			"Hello World",
		)
	}
}
