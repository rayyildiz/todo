package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"go.rayyildiz.dev/todo/pkg/adapter/handler/graphql"
	"go.rayyildiz.dev/todo/pkg/adapter/repository"
	"go.rayyildiz.dev/todo/pkg/adapter/service"
	"go.rayyildiz.dev/todo/pkg/port"
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
	r.Use(middleware.Heartbeat("/health"))

	repo := newRepository()
	todoSvc := service.NewTodoService(repo)

	rootResolver := graphql.NewGraphqlResolver(todoSvc)
	graphqlServer := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: rootResolver}))

	if os.Getenv("GRAPHQL_ENABLE_PLAYGROUND") == "true" {
		r.Handle("/api/docs", playground.Handler("GraphQL playground", "/api/query"))
	}
	r.Handle("/api/query", graphqlServer)

	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "4000"
	}
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("could not start server at :%s, %v", port, err)
	}
}

func newRepository() port.Repository {
	collection := os.Getenv("DOCSTORE_COLLECTION")
	if len(collection) < 5 { // mongo://, firestore://
		return repository.NewMemoryRepository()
	}

	coll, err := docstore.OpenCollection(context.Background(), collection)
	if err != nil {
		log.Fatalf("can't start with docstore, check your enviropment, %v", err)
	}
	return repository.NewDocstoreRepository(coll)
}
