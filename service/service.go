package service

import (
	"context"
	"go-api-todolist/models"
	"go-api-todolist/repository"
)

type Service interface {
	GetTask(ctx context.Context, id string) (*models.Todo, error)
	GetAllTasks(ctx context.Context) ([]*models.Todo, error)
	CreateTask(ctx context.Context, t models.Todo) (string, error)
	DeleteTask(ctx context.Context, id string) error
	UpdateTask(ctx context.Context, id string, t models.Todo) error
	Ping(ctx context.Context) error
}

type service_impl struct {
	mongodb repository.MongoDB
}

func New(m repository.MongoDB) Service {
	return &service_impl{
		mongodb: m,
	}
}
