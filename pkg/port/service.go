//go:generate mockgen -package=port -self_package=go.rayyildiz.dev/todo/pkg/port -destination=service_mock.go . Service
package port

import (
	"context"
	"go.rayyildiz.dev/todo/pkg/domain"
)

type ServiceUpdater interface {
	NewTodo(ctx context.Context, content string) (*domain.Todo, error)
	Delete(ctx context.Context, id string) error
	ToggleComplete(ctx context.Context, id string) error
}

type ServiceFinder interface {
	FindAll(ctx context.Context) ([]domain.Todo, error)
	FindById(ctx context.Context, id string) (*domain.Todo, error)
}

// Service represents the business layer.
type Service interface {
	ServiceFinder
	ServiceUpdater
}
