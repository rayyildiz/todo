package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.rayyildiz.dev/todo/pkg/port"
	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/memdocstore"
	"testing"
)

func TestDocstoreRepository_Store(t *testing.T) {
	ctx := context.Background()

	coll, err := docstore.OpenCollection(ctx, "mem://collection/id")
	require.NoError(t, err)
	defer coll.Close()

	var repo port.RepositorySaver = NewDocstoreRepository(coll)

	todo, err := repo.Store(ctx, "sample text")
	require.NoError(t, err)
	if assert.NotNil(t, todo) {
		assert.Equal(t, "sample text", todo.Content)
		assert.False(t, todo.Completed)
	}
}

func TestDocstoreRepository_FindAll(t *testing.T) {
	ctx := context.Background()

	coll, err := docstore.OpenCollection(ctx, "mem://collection/id")
	require.NoError(t, err)
	defer coll.Close()

	var repo port.Repository = NewDocstoreRepository(coll)
	repo.Store(ctx, "hello")
	repo.Store(ctx, "world")

	todos, err := repo.FindAll(ctx)
	require.NoError(t, err)
	if assert.NotEmpty(t, todos) {
		assert.Equal(t, 2, len(todos))
	}

}

func TestDocstoreRepository_FindById(t *testing.T) {
	ctx := context.Background()

	coll, err := docstore.OpenCollection(ctx, "mem://collection/id")
	require.NoError(t, err)
	defer coll.Close()

	var repo port.Repository = NewDocstoreRepository(coll)
	t1, _ := repo.Store(ctx, "text1")
	t2, _ := repo.Store(ctx, "text2")
	require.NotNil(t, t1)
	require.NotNil(t, t2)

	todo, err := repo.FindById(ctx, t1.ID)
	require.NoError(t, err)
	if assert.NotNil(t, todo) {
		assert.Equal(t, "text1", todo.Content)
	}
}

func TestDocstoreRepository_Toggle(t *testing.T) {
	ctx := context.Background()

	coll, err := docstore.OpenCollection(ctx, "mem://collection/id")
	require.NoError(t, err)
	defer coll.Close()

	var repo port.Repository = NewDocstoreRepository(coll)
	todo, err := repo.Store(ctx, "text1")
	require.NoError(t, err)
	actualTodo, err := repo.Toggle(ctx, todo.ID)
	require.NoError(t, err)
	if assert.NotNil(t, actualTodo) {
		assert.Equal(t, actualTodo.ID, todo.ID)
		assert.NotEqual(t, actualTodo.Completed, todo.Completed)
	}

}

func TestDocstoreRepository_Delete(t *testing.T) {
	ctx := context.Background()

	coll, err := docstore.OpenCollection(ctx, "mem://collection/id")
	require.NoError(t, err)
	defer coll.Close()

	var repo port.Repository = NewDocstoreRepository(coll)
	todo, err := repo.Store(ctx, "text1")

	err = repo.Delete(ctx, todo.ID)
	require.NoError(t, err)
}
