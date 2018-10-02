package server

import (
	"log"

	"projects/unrealengine.help/backend/db"
	"projects/unrealengine.help/backend/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// SrvData -
type SrvData struct {
	Port         string
	DataBaseName string
	URL          string
	User         string
	Password     string
	LocalHost    string
	DataBase     db.DataBase
	Echo         *echo.Echo
}

// SrvInterface -
type SrvInterface interface {
	Routes()
	Run()
}

// Create -
func Create(port, dataBasebName, url, user, password, localHost string) (SrvInterface, error) {
	e := echo.New()
	db := db.DataBase{}
	return &SrvData{Port: ":" + port, DataBaseName: dataBasebName, URL: url, User: user, Password: password, DataBase: db, Echo: e, LocalHost: localHost}, nil
}

// ConnectDB -
func (srv *SrvData) ConnectDB() *db.Store {
	store, err := srv.DataBase.Connect(srv.DataBaseName, srv.User, srv.URL, srv.Password)
	if err != nil {
		log.Printf("SERVER:-->::ConnectDB:::ERROR::NO CONNECTED IN DB: %s\n", err)
		panic(err)
	}
	return store
}

// Routes -
func (srv *SrvData) Routes() {
	srv.Echo.Static("/", "public")
	store := srv.ConnectDB()
	handlers := &handlers.Store{DB: store}

	// CORS
	srv.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{srv.LocalHost},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Logger Middleware
	srv.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// favicon
	srv.Echo.File("/favicon.ico", "public/favicon.ico")

	// Error path handler
	srv.Echo.HTTPErrorHandler = handlers.HTTPErrorRedirectHandler

	// API
	api := srv.Echo.Group("/api/v1")
	api.POST("/auth", handlers.Auth)
}

// Run -
func (srv *SrvData) Run() {
	srv.Echo.Logger.Fatal(srv.Echo.Start(srv.Port))
}
