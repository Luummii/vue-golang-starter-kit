package db

import (
	"database/sql"
	"vue-golang-starter-kit/backend/types"

	"github.com/labstack/gommon/log"
	// postgres drivers
	_ "github.com/lib/pq"
)

// DataBase for using DataBase
type DataBase struct {
}

var db *sql.DB

// Connect - соединение с БД
func Connect(database, user, url, password string) *DataBase {
	dsn := "postgres://" + user + ":" + password + "@" + url + ":5432/" + database

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &DataBase{}
}

// AddUser - Добавляем юзера в БД
func (*DataBase) AddUser(u types.User) {
	query := `INSERT INTO users (first, last, password, email) VALUES ($1, $2, $3, $4)`

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	r, err := stmt.Exec(u.First, u.Last, u.Password, u.Email)
	if err != nil {
		panic(err)
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		log.Error("Error insert in database")
	}
}
