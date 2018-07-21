package server

import (
	"github.com/labstack/echo"
)

// Server - a new server struct
type Server struct {
	Port   string
	Server *echo.Echo
}

// SInterface interface
type SInterface interface {
	Listen()
}

// Create - function creating a new server
func Create(port string) SInterface {
	return &Server{":" + port, echo.New()}
}

// Listen - function listeninig server in port
func (s *Server) Listen() {
	s.Server.Logger.Fatal(s.Server.Start(":5000"))
}
