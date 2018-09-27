package handler

import (
	"log"
	"net/http"
	"projects/vue-golang-starter-kit/backend/db"
)

// Store for using DataBase
type Store struct {
	DB *db.Store
}

// AddUsers bla bla bla
func (s *Store) AddUsers(w http.ResponseWriter, r *http.Request) {
	log.Printf("Auth: page")
}
