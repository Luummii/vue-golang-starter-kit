package db

import (
	"database/sql"

	"github.com/labstack/gommon/log"
	// postgres drivers
	_ "github.com/lib/pq"
)

type user struct {
	First    string
	Last     string
	Password string
	Email    string
}

var db *sql.DB

// Connect - DataBase connect
func Connect(database, user, url, password string) *sql.DB {
	dsn := "postgres://" + user + ":" + password + "@" + url + ":5432/" + database

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	addUser()

	return db
}

// AddUser - function add user in db
func addUser() {
	query := `INSERT INTO users (first, last, password, email) VALUES ($1, $2, $3, $4)`
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	u := user{"Andrey", "Zimin", "qwert", "asd@asd.asd"}
	r, err := stmt.Exec(u.First, u.Last, u.Password, u.Email)
	if err != nil {
		panic(err)
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		log.Error("Error insert in database")
	}
}
