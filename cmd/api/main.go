package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/felipedavid/go_crud/internal/models"

	_ "github.com/lib/pq"
)

type app struct {
	gopher *models.GopherModel
}

func main() {
	addr := *flag.String("addr", ":4000", "Address to listen on")
	dsn := *flag.String("dsn", "postgres://web:secret@localhost/go_crud?sslmode=disable", "Data Service Name")
	flag.Parse()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	a := app{
		gopher: &models.GopherModel{db},
	}

	log.Printf("Listening on %s\n", addr)
	err = http.ListenAndServe(addr, a.routes())
	log.Fatal(err)
}
