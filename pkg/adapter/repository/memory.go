package repository

import (
	"context"
	"github.com/google/uuid"
	"go.rayyildiz.dev/todo/pkg/domain"
	"go.rayyildiz.dev/todo/pkg/port"
)

type memoryRepository struct {
	container []domain.Todo
}

func NewMemoryRepository() *memoryRepository {
	return &memoryRepository{}
}
func (m *memoryRepository) FindById(ctx context.Context, id string) (*domain.Todo, error) {
	for _, con := range m.container {
		if con.ID == id {
			return &con, nil
		}
	}

	return nil, port.ErrRecordNotFound
}

func (m *memoryRepository) FindAll(ctx context.Context) ([]domain.Todo, error) {
	userId := port.UserFromContext(ctx)

	var models []domain.Todo

	for _, todo := range m.container {
		if userId == todo.User {
			models = append(models, todo)
		}
	}

	return models, nil
}

func (m *memoryRepository) Store(ctx context.Context, content string) (*domain.Todo, error) {
	model := domain.Todo{
		ID:        uuid.New().String(),
		Content:   content,
		Completed: false,
		User:      port.UserFromContext(ctx),
	}

	m.container = append(m.container, model)
	return &model, nil
}

func (m *memoryRepository) Toggle(ctx context.Context, id string) (*domain.Todo, error) {
	for i, con := range m.container {
		if con.ID == id {
			m.container[i].Completed = !m.container[i].Completed
			return &m.container[i], nil
		}
	}

	return nil, port.ErrRecordNotFound
}

func (m *memoryRepository) Delete(ctx context.Context, id string) error {
	for i, con := range m.container {
		if con.ID == id {
			m.container = append(m.container[:i], m.container[i+1:]...)
			return nil
		}
	}

	return port.ErrRecordNotFound
}
