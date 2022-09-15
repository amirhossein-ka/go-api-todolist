package service

import "go-api-todolist/repository"

type Service interface {

}

type service_impl struct {
    mongodb repository.MongoDB
}

func New(m repository.MongoDB) Service {
    return service_impl{
        mongodb: m,
    }
}
