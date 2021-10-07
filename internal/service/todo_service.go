package service

import (
	"playground/configuration_manager/internal/client"
	"playground/configuration_manager/internal/model"
)

type Service interface {
	CreateTodo(todo model.Todo) (model.Todo, error)
	Todos() []model.Todo
}


type service struct{
	Client client.Todo
}

func NewTodoService(todoClient client.Todo) Service {
	return &service{
		Client: todoClient,
	}
}

func (s service) CreateTodo(todo model.Todo) (model.Todo, error) {
	panic("implement me")
}

func (s service) Todos() []model.Todo {
	panic("implement me")
}
