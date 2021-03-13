package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"gocloud.dev/docstore"
	"log"
	"net/http"
	"os"
)

func init() {
	godotenv.Load()
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer)

	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "4000"
	}
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("could not start server at :%s, %v", port, err)
	}
}

func newDocStore(collection string) (*docstore.Collection, error) {
	if collection == "" {
		return nil, errors.New("$DOCSTORE_COLLECTION can't be nil")
	}

	return docstore.OpenCollection(context.Background(), collection)
}

func newDatabase(connStr string) (*sql.DB, error) {
	if connStr == "" {
		return nil, errors.New("please provide a db connection string")
	}

	// postgres://postgres:123456@localhost/postgres?sslmode=disable

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(4)

	err = db.Ping()
	return db, err
}
