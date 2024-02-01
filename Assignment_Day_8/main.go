package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var m = make(map[string]string)

type WebList struct {
	Websites []string `json:"urls"`
}

func checkStatus(web string) {
	resp, err := http.Get(web)
	if err != nil {
		m[web] = "DOWN"
		return
	}

	if resp.StatusCode == http.StatusOK {
		m[web] = "UP"
	} else {
		m[web] = "DOWN"
	}
}

func checkWebRespond(webs []string) {

	for i := 0; i < len(webs); i++ {
		// go checkStatus(webs[i])
		checkStatus(webs[i])
	}

	time.Sleep(1 * time.Second)
}

func handlePostWebStatus(w http.ResponseWriter, r *http.Request) {
	var web WebList

	err := json.NewDecoder(r.Body).Decode(&web)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	webList := web.Websites
	// go checkWebRespond(webList)
	checkWebRespond(webList)

}

func handlePrintingWebList(w http.ResponseWriter, r *http.Request) {
	_, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	// w.Write(res)
	io.WriteString(w, `{
			"urls":["https://www.google.com" , "https://www.saurabhpuri.com" , "https://www.facebook.com" , "https://joshsoftware.com"]
	}`)
}

func handleParticularWebsite(w http.ResponseWriter, r *http.Request) {
	/*
		// query := r.URL.Query()
		// filter := query.Get("name")
		// fmt.Println(filter)

		r.ParseForm()
		mapValue := r.Form
		// value, _ := json.Marshal(mapValue)

		// w.Write(value)

		protocol := "https://"
		filter := mapValue["name"][0]

		res, err := http.Get(protocol + filter)

		if err != nil {
			w.Write([]byte("DOWN"))
			return
		}

		if res.StatusCode == http.StatusOK {
			w.Write([]byte("UP"))
		} else {
			w.Write([]byte("DOWN"))
		}

		w.WriteHeader(http.StatusOK)
		// w.Write([]byte(filter))
	*/

	// we can get the value from the map also
	// r.ParseForm()
	// mapValue := r.Form
	// website := mapValue["name"][0]

	query := r.URL.Query()
	website := query.Get("name")

	filter := "https://" + website
	val, flag := m[filter]
	if flag {
		w.Write([]byte(val))
	} else {
		w.Write([]byte("Does't Exits in our databasae"))
	}

}

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/websites", handlePostWebStatus).Methods(http.MethodPost)
	mux.HandleFunc("/websites", handlePrintingWebList).Methods(http.MethodGet)
	// mux.HandleFunc("/websites", handlePrintingWebList)
	mux.HandleFunc("/website", handleParticularWebsite).Methods(http.MethodGet)
	fmt.Println("Starting Server...")
	http.ListenAndServe("localhost:9000", mux)

	go func() {
		for {
			for key := range m {
				checkStatus(key)
			}
			time.Sleep(1 * time.Minute)
		}
	}()
}
