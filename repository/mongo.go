package repository

import (
	"context"
	"go-api-todolist/models"
)

type (
	MongoDB interface {
		Create(ctx context.Context, t models.Todo) (any, error)
		ReadOne(ctx context.Context, id uint) (*models.Todo, error)
		ReadAll(ctx context.Context) ([]*models.Todo, error)
		Update(ctx context.Context, id uint, t *models.Todo) error
		Delete(ctx context.Context, id uint) error
	}
)
