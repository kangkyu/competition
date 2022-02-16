package main

import (
	"fmt"
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
	db.MustExec(`TRUNCATE users`)
	_, err = db.Exec(`INSERT INTO users (email) VALUES ($1)`, "lee@kangkyu.com")
	if err != nil {
		panic(err)
	}
	rows, err := db.Queryx(`SELECT * FROM users`)

	fmt.Println("Hello world")
	for rows.Next() {
		var u User
		err = rows.StructScan(&u)
		fmt.Printf("Hello, %v with id %d\n", u.Email, u.ID)
	}

	db.Close()
}
