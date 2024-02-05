package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	req, err := http.NewRequest("GET", "/websites", nil)

	if err != nil {
		t.Fatal(err) // it will print the error and stop the program
	}

	resp := httptest.NewRecorder()                     // record the response
	handler := http.HandlerFunc(handlePrintingWebList) // get the handler

	handler.ServeHTTP(resp, req)

	// check the status of code
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{
		"urls":["https://www.google.com" , "https://www.saurabhpuri.com" , "https://www.facebook.com" , "https://joshsoftware.com"]
	  }`

	if reflect.DeepEqual(resp.Body.String(), expected) {
		t.Errorf("Handler return the unexpectes response , got : %v and want %v ", resp.Body.String(), expected)
	}

}

func TestPostMethod(t *testing.T) {
	jsonData := []byte(`{"urls":["https://www.google.com" , "https://www.saurabhpuri.com" , "https://www.facebook.com" , "https://joshsoftware.com"] }`)

	// bytes.NewBuffer creates the new buffer and store the new data into it
	req, err := http.NewRequest("POST", "/websites", bytes.NewBuffer(jsonData))

	if err != nil {
		t.Fatal(err)
	}

	//  get the response
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handlePostWebStatus)
	handler.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Unexptected Output")
	}
}

func TestParticularWebsite(t *testing.T) {
	//For Valid Case
	req, err := http.NewRequest("GET", "/website?name=https://www.youtube.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handleParticularWebsite)

	handler.ServeHTTP(resp, req)

	// cheking from status code
	if resp.Code != http.StatusOK {
		t.Errorf("Unexpected error got %v , want %v", resp.Code, http.StatusOK)
	}

	expectedBody := "UP"

	// if body := resp.Body.String(); body != expectedBody {
	// 	t.Errorf("Expected body '%s', but got '%s'", expectedBody, body)
	// }
	if reflect.DeepEqual(resp.Body.String(), expectedBody) {
		t.Errorf("Handler return the unexpectes response , got : %v and want %v ", resp.Body.String(), expectedBody)
	}
}
