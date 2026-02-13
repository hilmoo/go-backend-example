package todo

import (
	"context"

	"github.com/hilmoo/go-backend-example/internal/model"
)

type DbRepository interface {
	Create(ctx context.Context, t *model.Todo) error
	GetByID(ctx context.Context, id string) (*model.Todo, error)
	Update(ctx context.Context, t *model.Todo) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, status string) ([]*model.Todo, error)
}
