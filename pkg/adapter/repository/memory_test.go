package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.rayyildiz.dev/todo/pkg/port"
	"testing"
)

func TestMemoryRepository_Store(t *testing.T) {
	ctx := context.Background()

	var repo port.RepositorySaver = NewMemoryRepository()

	todo, err := repo.Store(ctx, "hello")
	require.NoError(t, err)
	if assert.NotNil(t, todo) {
		assert.Equal(t, "hello", todo.Content)
		assert.False(t, todo.Completed)
	}
}

func TestMemoryRepository_FindAll(t *testing.T) {
	ctx := context.Background()

	var repo port.Repository = NewMemoryRepository()
	repo.Store(ctx, "do this")
	repo.Store(ctx, "do that")

	todos, err := repo.FindAll(ctx)
	require.NoError(t, err)
	if assert.NotEmpty(t, todos) {
		assert.Equal(t, 2, len(todos))
	}
}

func TestMemoryRepository_FindById(t *testing.T) {
	ctx := context.Background()

	var repo port.Repository = NewMemoryRepository()

	todo1, _ := repo.Store(ctx, "do")
	todo2, _ := repo.Store(ctx, "did")
	require.NotNil(t, todo1)
	require.NotNil(t, todo2)

	todo, err := repo.FindById(ctx, todo1.ID)
	require.NoError(t, err)
	if assert.NotNil(t, todo) {
		assert.Equal(t, "do", todo.Content)
	}
}

func TestMemoryRepository_Toggle(t *testing.T) {
	ctx := context.Background()

	var repo port.Repository = NewMemoryRepository()

	todo, err := repo.Store(ctx, "title")
	require.NoError(t, err)
	actualTodo, err := repo.Toggle(ctx, todo.ID)
	require.NoError(t, err)
	if assert.NotNil(t, actualTodo) {
		assert.Equal(t, actualTodo.ID, todo.ID)
		assert.NotEqual(t, actualTodo.Completed, todo.Completed)
	}

}
