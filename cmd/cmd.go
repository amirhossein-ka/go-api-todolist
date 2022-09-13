package cmd

import (
	"context"
	"go-api-todolist/config"
	"go-api-todolist/repository/mongo"
)

func Run(cfg *config.Config) error {
	mongodb, err := mongo.New(context.Background(), &cfg.Database)
	if err != nil {
		return err
	}

	_ = mongodb
	return nil
}
