package main

//need to access with the IP address
//get via hostname -I when SSH'd into the pi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handleIndex)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Running hello world on the pi")

	s.ListenAndServe()
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := make(map[string]string)

	resp["hello"] = "world"

	json.NewEncoder(w).Encode(resp)
}
