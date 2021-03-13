package port

import (
	"context"
	"go.rayyildiz.dev/todo/pkg/domain"
)

// RepositorySaver modifies or deletes data.
type RepositorySaver interface {
	Store(ctx context.Context, content string) (*domain.Todo, error)
	Toggle(ctx context.Context, id string) (*domain.Todo, error)
	Delete(ctx context.Context, id string) error
}

// RepositoryReader is read-only repo.
type RepositoryReader interface {
	FindById(ctx context.Context, id string) (*domain.Todo, error)
	FindAll(ctx context.Context) ([]domain.Todo, error)
}

type Repository interface {
	RepositoryReader
	RepositorySaver
}
