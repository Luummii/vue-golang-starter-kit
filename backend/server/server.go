package server

import (
	"vue-golang-starter-kit/backend/db"
	"vue-golang-starter-kit/backend/routes"

	"github.com/labstack/echo"
)

// Server - структура сервера
type Server struct {
	Port string
	DB   *db.DataBase
	Echo *echo.Echo
}

// SInterface специально умный интерфейс
type SInterface interface {
	Listen()
	CreateRoutes()
}

// Create - создаем новый сервер
func Create(port string, db *db.DataBase) SInterface {
	return &Server{":" + port, db, echo.New()}
}

// Listen - запускаем слушание на порту
func (srv *Server) Listen() {
	srv.Echo.Logger.Fatal(srv.Echo.Start(srv.Port))
}

// CreateRoutes - функция создания всех роутов и мидлов
func (srv *Server) CreateRoutes() {
	//usersGroup := srv.Echo.Group("/users")

	routes.Home(srv.Echo)
	//routes.Users(usersGroup)
}
