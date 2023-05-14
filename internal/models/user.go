package models

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID         int
	Username   string `sql:"not null;unique"`
	Token      string `sql:"not null;unique"`
	Count      int    `sql:"not null;default:0"`
	Status     int    `sql:"not null;default:1"`
	CreateTime time.Time
	UpdateTime time.Time
}

func InitDB(ctx context.Context) {
	db, err := sql.Open("sqlite3", "chat.db")
	if err != nil {
		panic(err)
	}
	go func() {
		<-ctx.Done()
		db.Close()
	}()
}
