// building a basic web app with gorilla
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	listenAddr string = ":8080"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var users []*User

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/content/{x}", handleContent).Methods("GET")
	r.HandleFunc("/user", handleCreateUser).Methods("POST")
	r.HandleFunc("/user", handleGetUser).Methods("GET")

	//create the server
	s := &http.Server{
		Addr:    listenAddr,
		Handler: r,
	}

	fmt.Println("Running basic web app")

	//run the server
	s.ListenAndServe()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	//make some demo data
	resp := make(map[string]string)

	resp["hello"] = "world"

	//set content type as JSON
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)

}

func handleContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resp := make(map[string]string)

	resp["content"] = vars["x"]

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	u := User{}

	json.NewDecoder(r.Body).Decode(&u)

	users = append(users, &u)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}
