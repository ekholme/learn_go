//goal is to build a demo-ish app that has a cacheing layer
//so we don't need to hit the database with each request
// see https://www.youtube.com/watch?v=7zDl-aPW9sg&t=1239s

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Id       int
	Username string
}

type Server struct {
	db    map[int]*User
	cache map[int]*User
	srv   *http.Server
}

func NewServer() *Server {

	db := make(map[int]*User)

	for i := 1; i <= 100; i++ {
		db[i] = &User{
			Id:       i,
			Username: fmt.Sprintf("user_%d", i),
		}
	}

	srv := &http.Server{
		Addr: ":8080",
	}

	return &Server{
		db:  db,
		srv: srv,
	}
}

// RESUME HERE
func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Fatal(err)
	}

	user, ok := s.db[id]
	if !ok {
		panic("user not found")
	}

	json.NewEncoder(w).Encode(user)
}
