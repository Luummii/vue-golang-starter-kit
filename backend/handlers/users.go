package handlers

import (
	"net/http"
	"vue-golang-starter-kit/backend/db"
	"vue-golang-starter-kit/backend/types"

	"github.com/labstack/echo"
)

// Store for using DataBase
type Store struct {
	DB *db.Store
}

// AddUsers bla bla bla
func (s *Store) AddUsers(c echo.Context) error {
	u := types.User{First: "Andrey", Last: "Zimin", Password: "golang", Email: "go@go.go"}
	s.DB.Add(u)
	return c.String(http.StatusOK, "!! horay you are on the secret amdin main page!")
}

// GetUsers bla bla bla
func (s *Store) GetUsers(c echo.Context) error {
	// ТУТ ЧИТАЕМ ИЗ БД
	return c.String(http.StatusOK, "horay you are on the secret amdin main page!")
}
