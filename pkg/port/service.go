package port

import (
	"context"
	"go.rayyildiz.dev/todo/pkg/domain"
)

// Service represents the business layer.
type Service interface {
	FindAll(ctx context.Context) ([]domain.Todo, error)
	NewTodo(ctx context.Context, content string) (*domain.Todo, error)
	Delete(ctx context.Context, id string) error
	ToggleComplete(ctx context.Context, id string) error
}
