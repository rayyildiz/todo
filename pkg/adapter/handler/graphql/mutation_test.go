package graphql

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.rayyildiz.dev/todo/pkg/domain"
	"go.rayyildiz.dev/todo/pkg/port"
	"testing"
)

func TestMutationResolver_New(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := port.NewMockService(ctrl)
	svc.EXPECT().NewTodo(ctx, "test").Return(&domain.Todo{
		ID:        "1",
		Content:   "test",
		Completed: false,
	}, nil)

	resolver := newMutationResolver(svc)

	todo, err := resolver.New(ctx, "test")
	require.NoError(t, err)
	if assert.NotNil(t, todo) {
		assert.Equal(t, "1", todo.ID)
		assert.Equal(t, false, todo.Completed)
	}
}

func TestMutationResolver_Toggle(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := port.NewMockService(ctrl)
	svc.EXPECT().ToggleComplete(ctx, "1").Return(nil)

	resolver := newMutationResolver(svc)

	b, err := resolver.Toggle(ctx, "1")
	if assert.NoError(t, err) {
		assert.Equal(t, "1", b)
	}
}

func TestMutationResolver_Delete(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := port.NewMockService(ctrl)
	svc.EXPECT().Delete(ctx, "1").Return(nil)

	resolver := newMutationResolver(svc)

	b, err := resolver.Delete(ctx, "1")
	if assert.NoError(t, err) {
		assert.True(t, b)
	}
}
