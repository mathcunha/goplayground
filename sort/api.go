package main

import (
	"encoding/json"
	"fmt"
	"github.com/mathcunha/goplayground/sort/msort"
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
	http.HandleFunc(ApiUrlPrefix+"/sort/bubble", func(w http.ResponseWriter, r *http.Request) {
		SortHandler(w, r, new(msort.BubbleSort))
	})
}

// DefaultHandler, the web server
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set(contentType, jsonContentType)
	w.WriteHeader(http.StatusOK)
}

// SortHandler, reads the body, loads json, calls sort function and writes the json
func SortHandler(w http.ResponseWriter, r *http.Request, s msort.Sort) {
	DefaultHandler(w, r)
	unordered, err := loadArray(r)
	if err != nil {
		fmt.Fprintf(w, "{\"error\":\"request body not valid\"}")
	} else {
		sort(w, unordered, s)
	}
}

// ShowOptions shows available sort algorithms
func ShowOptions(w http.ResponseWriter) {
	fmt.Fprintf(w, "{ \"algorithms\" : [{\"route\":\"%s\", \"algorithm\":\"%s\"}]}", ApiUrlPrefix+"/sort"+"/bubble", "Bubble Sort")
}

func loadArray(r *http.Request) (a []int, err error) {
	var body struct {
		Array []int `json:"array,required"`
	}
	d := json.NewDecoder(r.Body)
	err = d.Decode(&body)
	return body.Array, err
}

// sort calls sort function and writes the json to w
func sort(w http.ResponseWriter, u []int, s msort.Sort) (err error) {
	encoder := json.NewEncoder(w)
	s.Sort(&u)
	var body = struct {
		Ordered []int `json:"ordered,required"`
	}{u}
	err = encoder.Encode(body)
	return
}

func main() {
	SetupApiRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
