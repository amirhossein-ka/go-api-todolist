package service

import (
	"context"
	"go-api-todolist/models"
)

func (s *service_impl) GetTask(ctx context.Context, id string) (*models.Todo, error) {
	return s.mongodb.ReadOne(ctx, id)
}

func (s *service_impl) GetAllTasks(ctx context.Context) ([]*models.Todo, error) {
	return s.mongodb.ReadAll(ctx)
}

func (s *service_impl) CreateTask(ctx context.Context, t models.Todo) (string, error) {
	id, err := s.mongodb.Create(ctx, t)
	return id, err
}

func (s *service_impl) DeleteTask(ctx context.Context, id string) error {
	return s.mongodb.Delete(ctx, id)
}

func (s *service_impl) UpdateTask(
	ctx context.Context,
	id string,
	t models.Todo,
) error {
	return s.mongodb.Update(ctx, id, t)
}

func (s *service_impl) Ping(ctx context.Context) error {
	return s.mongodb.Ping(ctx)
}
