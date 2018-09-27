package main

import (
	"flag"
	"log"
	"os"
	"projects/vue-golang-starter-kit/backend/server"

	"github.com/joho/godotenv"
)

var port string
var dataBasebName string
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
}

func main() {
	s, err := server.Create(port, dataBasebName, url, user, password)
	if err != nil {
		log.Printf("MAIN:-->::main::ERROR::NO CREATE SERVER: %s\n", err)
		panic(err)
	}
	s.Routes()
	s.Run()
}
