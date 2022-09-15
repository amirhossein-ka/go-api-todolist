package cmd

import (
	"context"
	"go-api-todolist/config"
	"go-api-todolist/controller/mux"
	"go-api-todolist/repository/mongo"
	"go-api-todolist/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) error {
	mongodb, err := mongo.New(context.Background(), &cfg.Database)
	if err != nil {
		return err
	}

	srv := service.New(mongodb)

	controller := mux.New(srv)

	go func() {
		err := controller.Start(":8000")
		if err != nil {
			log.Fatal(err)
		}
	}()

	//gracefully stop server

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	if err := controller.Stop(); err != nil {
		log.Fatal(err)
	}
	return nil
}
