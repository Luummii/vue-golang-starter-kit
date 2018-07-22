package db

import (
	"database/sql"
	"vue-golang-starter-kit/backend/types"

	"github.com/labstack/gommon/log"
	// postgres drivers
	_ "github.com/lib/pq"
)

// DataBase бла бла бла
type DataBase struct {
}

// Store бла бла бла
type Store struct {
	DB *sql.DB
}

// Connect - бла бла бла
func (*DataBase) Connect(database, user, url, password string) (*Store, error) {
	dsn := "postgres://" + user + ":" + password + "@" + url + ":5432/" + database

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &Store{DB: db}, nil
}

// Add - бла бла бла
func (s *Store) Add(u types.User) {
	query := `INSERT INTO users (first, last, password, email) VALUES ($1, $2, $3, $4)`

	stmt, err := s.DB.Prepare(query)
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
