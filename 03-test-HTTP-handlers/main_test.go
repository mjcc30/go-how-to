package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// test handleGetFoo function with recorder
func TestHandleGetFooRR(t *testing.T) {
	rr := httptest.NewRecorder() // create a new recorder
	req, err := http.NewRequest(http.MethodGet, "", nil) // create a new request
	if err != nil {
		t.Error(err)
	}

	handleGetFoo(rr, req) // call handleGetFoo function

	resp := rr.Result() // get the response

	checkResponse(resp, t) // check response
}

// test handleGetFoo function
func TestHandleGetFoo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGetFoo)) // create a test server
	resp, err := http.Get(server.URL) // send a GET request to the test server
	if err != nil {
		t.Error(err)
	}

	checkResponse(resp, t) // check response
}

func checkResponse(resp *http.Response, t *testing.T)  {
	if resp.StatusCode != http.StatusOK { // check the status code
		t.Errorf("expected 200 but get %d", resp.StatusCode)
	}
	defer resp.Body.Close() // close the response body

	expected := "FOO"
	b, err := ioutil.ReadAll(resp.Body) // read the response body
	if err != nil {
		t.Error(err)
	}

	if string(b) != expected { // check the response body
		t.Errorf("expected %s but we got %s", expected, string(b))
	}
}
