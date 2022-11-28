// building a basic web app with gorilla
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	listenAddr string = ":8080"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")

	//create the server
	s := &http.Server{
		Addr:    listenAddr,
		Handler: r,
	}

	//run the server
	s.ListenAndServe()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	//make some demo data
	resp := make(map[string]string)

	resp["hello"] = "world"

	jResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("couldn't marshal json data: %s", err)
	}

	//set status
	w.WriteHeader(http.StatusOK)

	//set content type as JSON
	w.Header().Add("Content-Type", "application/json")

	//write this out as json
	w.Write(jResp)
	//this is currently showing up as plain text, but we can fix this later
}
