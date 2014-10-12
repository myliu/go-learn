package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
)

func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://mliu@localhost/lesson4?sslmode=disable")
	PanicIf(err)
	return db
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	m := martini.Classic()
	m.Map(SetupDB())

	m.Get("/", func(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
		rows, err := db.Query("SELECT title FROM books")
		PanicIf(err)
		defer rows.Close()

		for rows.Next() {
			var title string
			if err := rows.Scan(&title); err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(rw, "Title: %s\n", title)
		}
	})

	m.Run()
}
