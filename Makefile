graphql:
	go run github.com/99designs/gqlgen

generate:
	go generate  ./...

test:
	go test -cover ./...
