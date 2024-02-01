package main

import (
	_ "fmt"
	"net/http"
	"net/http/httptest"
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

	// expected := `{
	// 	"urls":["https://www.google.com" , "https://www.saurabhpuri.com" , "https://www.facebook.com" , "https://joshsoftware.com"]
	//   }`

	// wantValue := `{
	// 	"urls":["https://www.google.com" , "https://www.saurabhpuri.com" , "https://www.facebook.com" , "https://joshsoftware.com"]
	//   }`

	// fmt.Printf("resp--------------> %#v\n", resp.Body.String())
	// fmt.Printf("The type : %T", resp.Body.String())

	// if resp.Body.String() != expected {
	// 	t.Errorf("Handler return the unexpectes response , got : %v and want %v ", resp.Body.String(), expected)
	// }

	respCode := resp.Code
	givenCode := http.StatusOK

	if respCode != givenCode {
		t.Errorf("Handler return the unexpectes response , got : %v and want %v ", respCode, givenCode)
	}

}
