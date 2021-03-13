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

func TestResolver_Query(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := port.NewMockService(ctrl)
	svc.EXPECT().FindAll(ctx).Return([]domain.Todo{
		{
			ID:        "1",
			Content:   "1",
			Completed: false,
		},
		{
			ID:        "2",
			Content:   "completed",
			Completed: true,
		},
		{
			ID:        "3",
			Content:   "third",
			Completed: false,
		},
	}, nil)

	query := newQueryResolver(svc)

	todos, err := query.Todos(ctx)
	require.NoError(t, err)
	if assert.NotNil(t, todos) {
		assert.Equal(t, 3, len(todos))
	}
}
