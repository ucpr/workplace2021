package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

type Room struct {
	Id      int64  `db:"room_id, primarykey, autoincrement"`
	Title   string `db:"title"`
	Message string `db:"message"`
}

func main() {
	db, err := sql.Open("postgres", "user=root password=root dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln("Fail: %v", err)
		return
	}
	dbmap := &gorp.DbMap{
		Db:      db,
		Dialect: gorp.PostgresDialect{},
	}
	defer dbmap.Db.Close()

	dbmap.AddTableWithName(Room{}, "rooms").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatalln("Create tables failed: %v", err)
	}
}
