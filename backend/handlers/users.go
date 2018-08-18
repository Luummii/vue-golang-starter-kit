package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"projects/vue-golang-starter-kit/backend/db"
	"projects/vue-golang-starter-kit/backend/types"

	"github.com/labstack/echo"
)

// Store for using DataBase
type Store struct {
	DB *db.Store
}

// AddUsers bla bla bla
func (s *Store) AddUsers(c echo.Context) error {
	defer c.Request().Body.Close()

	u := types.User{}

	body := json.NewDecoder(c.Request().Body)
	err := body.Decode(&u)

	if err != nil {
		log.Printf("Failed processing AddUsers request: %s\n", err)
		return c.String(http.StatusBadRequest, "Ups!")
	}

	s.DB.Add(u)
	return c.String(http.StatusOK, "Horay!")
}

// GetUsers bla bla bla
func (s *Store) GetUsers(c echo.Context) error {
	// ТУТ ЧИТАЕМ ИЗ БД
	return c.String(http.StatusOK, "бла бла бла")
}
