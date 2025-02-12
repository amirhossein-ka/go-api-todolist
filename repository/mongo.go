package repository

import (
	"context"
	"go-api-todolist/models"
)

type (
	MongoDB interface {
		Create(ctx context.Context, t models.Todo) (string, error)
		ReadOne(ctx context.Context, id string) (*models.Todo, error)
		ReadAll(ctx context.Context) ([]*models.Todo, error)
		Update(ctx context.Context, id string, t models.Todo) error
		Delete(ctx context.Context, id string) error
		Ping(ctx context.Context) error
	}
)
