package main

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	Email       string    `db:"email"`
	ID          int64     `db:"id"`
	TimeCreated time.Time `db:"time_created"`
}

func main() {
	db, err := sqlx.Open("postgres", "postgresql://ugp:abcd1234@localhost:5432/competition_development?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	

	db.Close()
}
