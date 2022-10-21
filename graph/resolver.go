package graph

import "github.com/hxrxchang/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type IResolver interface {
	FindTodo(string) *model.Todo
	FindUser(string) *model.User
}

type Resolver struct{
	todos []*model.Todo
	users []*model.User
}

func (r *Resolver) FindUser(ID string) *model.User {
	var user *model.User
	users := r.users
	for _, v := range users {
		if v.ID == ID {
			user = v
			break
		}
	}
	return user
}

func (r *Resolver) FindTodo(ID string) *model.Todo {
	var todo *model.Todo
	todos := r.todos
	for _, v := range todos {
		if v.ID == ID {
			todo = v
			break
		}
	}
	return todo
}
