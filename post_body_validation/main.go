package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

const (
	listenAddr string = ":8080"
)

type User struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		msg := "hello, world"

		w.WriteHeader(http.StatusOK)

		w.Write([]byte(msg))
	case http.MethodPost:
		var u *User

		validate := validator.New()

		json.NewDecoder(r.Body).Decode(&u)

		err := validate.Struct(u)

		if err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		msg := "hello, " + u.FirstName + " " + u.LastName

		w.WriteHeader(http.StatusOK)

		w.Write([]byte(msg))
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", handleIndex)

	fmt.Println("Running server")

	err := http.ListenAndServe(listenAddr, r)

	log.Fatal(err)
}
