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

	// slc := []string{""}

	// op, _ := json.Marshal(slc)
	expected := `{
		"urls":["https://www.google.com" , "https://www.saurabhpuri.com" , "https://www.facebook.com" , "https://joshsoftware.com"]
	  }`

	// wantValue := `{
	// 	"urls":["https://www.google.com" , "https://www.saurabhpuri.com" , "https://www.facebook.com" , "https://joshsoftware.com"]
	//   }`

	// fmt.Printf("resp--------------> %#v\n", resp.Body.String())
	// fmt.Printf("The type : %T", resp.Body.String())

	if reflect.DeepEqual(resp.Body.String(), expected) {
		t.Errorf("Handler return the unexpectes response , got : %v and want %v ", resp.Body.String(), expected)

	}

	// respCode := resp.Code
	// givenCode := http.StatusOK

	// if respCode != givenCode {
	// 	t.Errorf("Handler return the unexpectes response , got : %v and want %v ", respCode, givenCode)
	// }

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

	// expectedOutpput := `{"urls":["https://www.google.com" , "https://www.saurabhpuri.com" , "https://www.facebook.com" , "https://joshsoftware.com"] }`

	// fmt.Println("The value :", resp.Body.String())
	// if resp.Body.String() != expectedOutpput {
	// 	t.Errorf("Unexpected outpput got : %v and want : %v\n", resp.Body.String(), expectedOutpput)
	// }

	if resp.Code != http.StatusOK {
		t.Errorf("Unexptected Output")
	}
}
