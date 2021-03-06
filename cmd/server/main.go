package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	driver "go.rayyildiz.dev/todo/pkg/adapter/handler/graphql"
	driven "go.rayyildiz.dev/todo/pkg/adapter/repository"
	"go.rayyildiz.dev/todo/pkg/adapter/service"
	"go.rayyildiz.dev/todo/pkg/port"
	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/gcpfirestore"
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

	router.Use(versionHandler)
	router.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer, middleware.Heartbeat("/health"))
	router.Use(contextHandler)

	repo := newRepository()
	todoSvc := service.NewTodoService(repo)

	rootResolver := driver.NewGraphqlResolver(todoSvc)
	graphqlServer := handler.NewDefaultServer(driver.NewExecutableSchema(driver.Config{Resolvers: rootResolver}))

	if os.Getenv("GRAPHQL_ENABLE_PLAYGROUND") == "true" {
		router.Handle("/api/docs", playground.Handler("GraphQL playground", "/api/query"))
	}

	router.Handle("/api/query", graphqlServer)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}).Handler(router)

	serverPort := os.Getenv("PORT")
	if len(serverPort) < 1 {
		serverPort = "4000"
	}
	log.Printf("server startinng :%s , version: %s", serverPort, version)
	if err := http.ListenAndServe(":"+serverPort, handler); err != nil {
		log.Fatalf("could not start server at :%s, %v", serverPort, err)
	}
}

func versionHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("x-api-version", version)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func contextHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := port.ContextWithUser(r.Context(), r.Header.Get("Authorization"))
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func newRepository() port.Repository {
	collection := os.Getenv("DOCSTORE_COLLECTION")
	if len(collection) < 5 { // mongo://, firestore://
		return driven.NewMemoryRepository()
	}

	coll, err := docstore.OpenCollection(context.Background(), collection)
	if err != nil {
		log.Fatalf("can't start with docstore, check your enviropment, %v", err)
	}
	return driven.NewDocstoreRepository(coll)
}
