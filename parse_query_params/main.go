package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func DemoGet(w http.ResponseWriter, r *http.Request) {

	urlParams := r.URL.Query()

	//in a use case for another project, i want the urlParams to be encoded as a map[string]string
	//but this won't always be the use case
	m := make(map[string]string, len(urlParams))

	for i, v := range urlParams {
		//in prod, we probably don't want to error out
		if len(v) > 1 {
			log.Fatal("query parameters should all be length 1")
		}
		s := strings.Join(v, "")
		m[i] = s
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(m)
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/demoget", DemoGet)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Running demo program")

	s.ListenAndServe()
}
