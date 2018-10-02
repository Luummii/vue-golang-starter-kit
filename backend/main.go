package main

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	"projects/unrealengine.help/backend/server"
)

var port string
var dataBasebName string
var localHost string
var user string
var url string
var password string

func init() {
	flag.StringVar(&port, "port", "5000", "Assigning the port that the server should listen on")
	flag.Parse()

	if err := godotenv.Load("config.env"); err != nil {
		panic(err)
	}

	if _port := os.Getenv("PORT"); len(_port) > 0 {
		port = _port
	}

	if _dbName := os.Getenv("DATABASENAME"); len(_dbName) > 0 {
		dataBasebName = _dbName
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

	if _localHost := os.Getenv("LOCALHOST"); len(_localHost) > 0 {
		localHost = _localHost
	}
}

func main() {
	srv, err := server.Create(port, dataBasebName, url, user, password, localHost)
	if err != nil {
		panic(err)
	}
	srv.Routes()
	srv.Run()
}
