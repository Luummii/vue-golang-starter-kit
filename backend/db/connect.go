package db

import (
	"database/sql"

	"log"
	// this is postgres drivers
	_ "github.com/lib/pq"
)

// DataBase -
type DataBase struct {
}

// Store -
type Store struct {
	DB *sql.DB
}

// Connect -
func (*DataBase) Connect(database, user, url, password string) (*Store, error) {
	dsn := "postgres://" + user + ":" + password + "@" + url + ":5432/" + database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("DB:-->::Connect::ERROR DB OPEN:%s\n", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Printf("DB:-->::Connect::ERROR DB PING:%s\n", err)
		panic(err)
	}

	return &Store{DB: db}, nil
}
