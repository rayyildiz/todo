package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.rayyildiz.dev/todo/pkg/domain"
	"go.rayyildiz.dev/todo/pkg/port"
	"gocloud.dev/docstore"
	"io"
)

const maxLimit = 100

type docstoreRepository struct {
	collection *docstore.Collection
}

func NewDocstoreRepository(coll *docstore.Collection) *docstoreRepository {
	return &docstoreRepository{collection: coll}
}

func (r docstoreRepository) FindById(ctx context.Context, id string) (*domain.Todo, error) {
	model := domain.Todo{ID: id}
	if err := r.collection.Get(ctx, &model); err != nil {
		return nil, fmt.Errorf("while getting by id , %w", err)
	}

	return &model, nil
}

func (r docstoreRepository) FindAll(ctx context.Context) ([]domain.Todo, error) {
	var models []domain.Todo

	userId := port.UserFromContext(ctx)

	it := r.collection.Query().Limit(maxLimit).Get(ctx)
	for {
		model := domain.Todo{}
		err := it.Next(ctx, &model)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("while getting all models, %w", err)
		}
		if model.User == userId {
			models = append(models, model)
		}
	}

	return models, nil
}

func (r docstoreRepository) Store(ctx context.Context, content string) (*domain.Todo, error) {
	model := domain.Todo{
		ID:        uuid.New().String(),
		Content:   content,
		Completed: false,
		User:      port.UserFromContext(ctx),
	}
	err := r.collection.Create(ctx, &model)
	if err != nil {
		return nil, fmt.Errorf("while storing model, %w", err)
	}
	return &model, err
}

func (r docstoreRepository) Toggle(ctx context.Context, id string) (*domain.Todo, error) {
	todo, err := r.FindById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("while toogle status, %w", err)
	}
	todo.Completed = !todo.Completed
	err = r.collection.Update(ctx, todo, docstore.Mods{"completed": todo.Completed})
	if err != nil {
		return nil, fmt.Errorf("while updateing status, %w", err)
	}

	return todo, nil
}

func (r docstoreRepository) Delete(ctx context.Context, id string) error {
	err := r.collection.Delete(ctx, &domain.Todo{ID: id})
	if err != nil {
		return fmt.Errorf("while deleting model, %w", err)
	}
	return nil
}
