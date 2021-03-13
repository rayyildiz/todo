package graphql

import (
	context "context"
	"errors"
	"go.rayyildiz.dev/todo/pkg/port"
	"log"
)

type queryResolver struct {
	service port.Service
}

func newQueryResolver(svc port.Service) QueryResolver {
	return &queryResolver{svc}
}

func (q queryResolver) Todos(ctx context.Context) ([]*Todo, error) {

	models, err := q.service.FindAll(ctx)
	if err != nil {
		log.Printf("service error, %v", err)
		return nil, errors.New("could not get todo list")
	}

	var todos []*Todo
	for _, model := range models {
		todos = append(todos, &Todo{
			ID:        model.ID,
			Content:   model.Content,
			Completed: model.Completed,
		})
	}

	return todos, nil
}
