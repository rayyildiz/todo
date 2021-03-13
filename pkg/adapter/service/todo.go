package service

import (
	"context"
	"fmt"
	"go.rayyildiz.dev/todo/pkg/domain"
	"go.rayyildiz.dev/todo/pkg/port"
	"log"
)

type todoService struct {
	repo port.Repository
}

func NewTodoService(repo port.Repository) *todoService {
	return &todoService{repo}
}

func (t todoService) FindAll(ctx context.Context) ([]domain.Todo, error) {
	todos, err := t.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not find any todo, %w", err)
	}

	return todos, nil
}

func (t todoService) NewTodo(ctx context.Context, content string) (*domain.Todo, error) {
	todo, err := t.repo.Store(ctx, content)
	if err != nil {
		return nil, fmt.Errorf("todo service, %w", err)
	}
	log.Printf("inserted a new todo, %#v", todo)

	return todo, nil
}

func (t todoService) Delete(ctx context.Context, id string) error {
	err := t.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("service, %w", err)
	}
	return nil
}

func (t todoService) ToggleComplete(ctx context.Context, id string) error {
	todo, err := t.repo.Toggle(ctx, id)
	if err != nil {
		return fmt.Errorf("service, %w", err)
	}
	log.Printf("todo update %#v", todo)
	return nil
}
