package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.rayyildiz.dev/todo/pkg/domain"
	"go.rayyildiz.dev/todo/pkg/port"
	"testing"
)

func TestTodoService_NewTodo(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := port.NewMockRepository(ctrl)
	repo.EXPECT().Store(ctx, "hello").Return(&domain.Todo{
		ID:        "1",
		Content:   "hello",
		Completed: false,
	}, nil)

	svc := NewTodoService(repo)
	todo, err := svc.NewTodo(ctx, "hello")
	require.NoError(t, err)
	if assert.NotNil(t, todo) {
		assert.Equal(t, "1", todo.ID)
		assert.Equal(t, "hello", todo.Content)
		assert.Equal(t, false, todo.Completed)
	}
}

func TestTodoService_FindAll(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := port.NewMockRepository(ctrl)
	repo.EXPECT().FindAll(ctx).Return([]domain.Todo{
		{
			ID:        "1",
			Content:   "test1",
			Completed: false,
		},
		{
			ID:        "2",
			Content:   "test2",
			Completed: false,
		},
	}, nil)

	svc := NewTodoService(repo)

	todos, err := svc.FindAll(ctx)
	require.NoError(t, err)
	if assert.NotNil(t, todos) {
		assert.Equal(t, 2, len(todos))
		assert.Equal(t, "2", todos[1].ID)
		assert.Equal(t, "test1", todos[0].Content)
	}
}

func TestTodoService_Delete(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := port.NewMockRepository(ctrl)
	repo.EXPECT().Delete(ctx, "1").Return(nil)
	repo.EXPECT().Delete(ctx, "2").Return(port.ErrRecordNotFound)

	svc := NewTodoService(repo)
	err1 := svc.Delete(ctx, "1")
	err2 := svc.Delete(ctx, "2")

	require.NoError(t, err1)
	require.Errorf(t, err2, port.ErrRecordNotFound.Error())
}

func TestTodoService_ToggleComplete(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := port.NewMockRepository(ctrl)
	repo.EXPECT().Toggle(ctx, "1").Return(&domain.Todo{
		ID:        "1",
		Content:   "hello",
		Completed: false,
	}, nil)
	repo.EXPECT().Toggle(ctx, "2").Return(nil, port.ErrRecordNotFound)

	svc := NewTodoService(repo)

	err1 := svc.ToggleComplete(ctx, "1")
	err2 := svc.ToggleComplete(ctx, "2")

	require.NoError(t, err1)
	require.Errorf(t, err2, port.ErrRecordNotFound.Error())
}
