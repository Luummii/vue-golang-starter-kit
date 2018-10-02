package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"projects/unrealengine.help/backend/db"
)

// Store -
type Store struct {
	DB *db.Store
}

// Auth -
func (s *Store) Auth(c echo.Context) error {
	defer c.Request().Body.Close()

	log.Printf("Auth")
	return c.String(http.StatusOK, "Horay!")
}
