package routes

import (
	"github.com/labstack/echo"
)

// Home - получение главной страницы
func Home(e *echo.Echo) {
	e.Static("/", "frontend/dist")
}
