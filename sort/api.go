package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	ApiUrlPrefix    = "/api/v1"
	contentType     = "Content-Type"
	jsonContentType = "application/json;charset=UTF-8"
)

// SetupApiRoutes is used within main to initialize all of the routes
func SetupApiRoutes() {
	http.HandleFunc(ApiUrlPrefix+"/sort", func(w http.ResponseWriter, r *http.Request) {
		DefaultHandler(w, r)
		ShowOptions(w)
	})
}

// DefaultHandler, the web server
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set(contentType, jsonContentType)
	w.WriteHeader(http.StatusOK)
}

// ShowOptions shows available sort algorithms
func ShowOptions(w http.ResponseWriter) {
	fmt.Fprintf(w, "{ \"algorithms\" : [{\"route\":\"%s\", \"algorithm\":\"%s\"}]}", ApiUrlPrefix+"/sort"+"/bubble", "Bubble Sort")
}

func main() {
	SetupApiRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
