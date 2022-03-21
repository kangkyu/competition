package main

import (
	"os"
	"time"

	mypubsub "example.com/competition/pubsub"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	Email       string    `db:"email"`
	ID          int64     `db:"id"`
	TimeCreated time.Time `db:"time_created"`
}

const DATABASE_URL = "postgresql://ugp:abcd1234@localhost:5432/competition_development?sslmode=disable"

func main() {
	db, err := sqlx.Open("postgres", DATABASE_URL)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	err = mypubsub.Publish(os.Stdout, "votes-topic", "Hello!")
	if err != nil {
		panic(err)
	}

	db.Close()
}
