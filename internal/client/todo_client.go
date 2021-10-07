package client

import (
	"net/http"
	"playground/configuration_manager/internal/model"
)

type Todo interface {
	GetTodos() []model.Todo
	NewTodo(todo model.Todo) (model.Todo, error)
}

type Client struct {
	URL string
	httpClient *http.Client
}

func NewTodoClient(url string) Todo {
	client := http.Client{}
	return &Client{
		URL: url,
		httpClient: &client,
	}
}

func (c Client) GetTodos() []model.Todo {
	panic("implement me")
}

func (c Client) NewTodo(todo model.Todo) (model.Todo, error) {
	panic("implement me")
}

