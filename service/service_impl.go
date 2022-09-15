package service

import (
	"context"
	"go-api-todolist/models"
)

func (s *service_impl) GetTask(ctx context.Context, id uint) (*models.Todo, error) {
	return s.mongodb.ReadOne(ctx, id)
}

func (s *service_impl) GetAllTasks(ctx context.Context) ([]*models.Todo, error) {
	return s.mongodb.ReadAll(ctx)
}

func (s *service_impl) CreateTask(ctx context.Context, t models.Todo) error {
	_, err := s.mongodb.Create(ctx, t)
	return err
}

func (s *service_impl) DeleteTask(ctx context.Context, id uint) error {
	return s.mongodb.Delete(ctx, id)
}

func (s *service_impl) UpdateTask(
	ctx context.Context,
	id uint,
	t *models.Todo,
) error {
	return s.mongodb.Update(ctx, id, t)
}
