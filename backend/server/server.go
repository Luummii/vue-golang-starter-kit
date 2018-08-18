package server

import (
	"projects/vue-golang-starter-kit/backend/db"
	"projects/vue-golang-starter-kit/backend/handlers"

	"github.com/labstack/echo"
)

// CRV - конфиг глобальных данных
type CRV struct {
	Port         string
	DataBaseName string
	URL          string
	User         string
	Password     string
	Echo         *echo.Echo
	DataBase     db.DataBase
}

// SInterface бла бла бла
type SInterface interface {
	CreateRoutes()
	Run()
}

// Create - бла бла бла
func Create(port, dbName, url, user, password string) (SInterface, error) {
	e := echo.New()
	db := db.DataBase{}
	return &CRV{Port: ":" + port, DataBaseName: dbName, URL: url, User: user, Password: password, Echo: e, DataBase: db}, nil
}

// CreateRoutes - бла бла бла
func (srv *CRV) CreateRoutes() {
	usersGroup := srv.Echo.Group("/users")

	store, err := srv.DataBase.Connect(srv.DataBaseName, srv.User, srv.URL, srv.Password)
	handlers := &handlers.Store{DB: store}
	if err != nil {
		panic(err)
	}

	srv.Echo.Static("/", "frontend/dist")
	usersGroup.POST("/add", handlers.AddUsers)
}

// Run - бла бла бла
func (srv *CRV) Run() {
	srv.Echo.Logger.Fatal(srv.Echo.Start(srv.Port))
}
