package graphql

import (
	"context"
	"errors"
	"go.rayyildiz.dev/todo/pkg/port"
	"log"
)

type mutationResolver struct {
	service port.Service
}

func newMutationResolver(svc port.Service) *mutationResolver {
	return &mutationResolver{service: svc}
}

func (m mutationResolver) New(ctx context.Context, content string) (*Todo, error) {

	todo, err := m.service.NewTodo(ctx, content)
	if err != nil {
		log.Printf("while getting backend, %v", err)
		return nil, errors.New("can't create a new todo")
	}
	return &Todo{
		ID:        todo.ID,
		Content:   todo.Content,
		Completed: todo.Completed,
	}, nil
}

func (m mutationResolver) Toggle(ctx context.Context, id string) (bool, error) {
	err := m.service.ToggleComplete(ctx, id)
	if err != nil {
		log.Printf("backend service, %v", err)
		return false, errors.New("can't change status")
	}
	return true, nil
}

func (m mutationResolver) Delete(ctx context.Context, id string) (bool, error) {
	err := m.service.Delete(ctx, id)
	if err != nil {
		log.Printf("backend service, %v", err)
		return false, errors.New("can't delete todo")
	}
	return true, err
}
