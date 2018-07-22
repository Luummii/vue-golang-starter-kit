package main

import (
	"flag"
	"os"
	"vue-golang-starter-kit/backend/server"

	"github.com/joho/godotenv"
)

var port string
var dbName string
var user string
var url string
var password string

func init() {
	// для примера, можно запускать сервер с консоли с флагом:--> go run main.go -port=8080
	flag.StringVar(&port, "port", "5000", "Assigning the port that the server should listen on")
	flag.Parse()

	if err := godotenv.Load("config.env"); err != nil {
		panic(err)
	}

	if _port := os.Getenv("PORT"); len(_port) > 0 {
		port = _port
	}

	if _dbName := os.Getenv("DATABASENAME"); len(_dbName) > 0 {
		dbName = _dbName
	}

	if _user := os.Getenv("USER"); len(_user) > 0 {
		user = _user
	}

	if _url := os.Getenv("URL"); len(_url) > 0 {
		url = _url
	}

	if _password := os.Getenv("PASSWORD"); len(_password) > 0 {
		password = _password
	}
}

func main() {
	s, err := server.Create(port, dbName, url, user, password)
	if err != nil {
		panic(err)
	}

	s.CreateRoutes()
	s.Run()
}
