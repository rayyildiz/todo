package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"go.rayyildiz.dev/todo/pkg/domain"
	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/gcpfirestore"
	"io"
	"log"
	"net/http"
	"os"
)

var version string // do not delete or modify

func init() {
	godotenv.Load()
}

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer, middleware.Heartbeat("/health"))

	coll := newRepository()

	router.Post("/clean", func(writer http.ResponseWriter, request *http.Request) {

		err := deleteAll(request.Context(), coll)
		if err != nil {
			log.Printf("error while deleteing documents, %v", err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusAccepted)
		writer.Write([]byte("done"))
	})

	serverPort := os.Getenv("PORT")
	if len(serverPort) < 1 {
		serverPort = "4000"
	}
	log.Printf("server startinng :%s , version: %s", serverPort, version)
	if err := http.ListenAndServe(":"+serverPort, router); err != nil {
		log.Fatalf("could not start server at :%s, %v", serverPort, err)
	}
}

func deleteAll(ctx context.Context, coll *docstore.Collection) error {
	const limit = 1000

	it := coll.Query().Limit(limit).Get(ctx)
	defer it.Stop()

	var todos []domain.Todo

	for {
		var doc domain.Todo
		err := it.Next(ctx, &doc)
		if err == io.EOF {
			log.Printf("end of list")
			break
		}
		if err != nil {
			return fmt.Errorf("while getting document, %w", err)
		}
		todos = append(todos, doc)
	}

	log.Printf("found %dd items", len(todos))

	list := coll.Actions()
	for _, doc := range todos {
		list.Delete(&doc)
	}

	err := list.Do(ctx)

	if err != nil {
		return fmt.Errorf("cant run actions list %w", err)
	}
	return nil
}

func newRepository() *docstore.Collection {
	collection := os.Getenv("DOCSTORE_COLLECTION")
	if len(collection) < 5 { // mongo://, firestore://
		log.Fatalf("it has to be an collection")
	}

	coll, err := docstore.OpenCollection(context.Background(), collection)
	if err != nil {
		log.Fatalf("can't start with docstore, check your enviropment, %v", err)
	}
	return coll
}
