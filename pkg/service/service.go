package service

import (
	todoapp "todo-app"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todoapp.User) (int,error)
	GenerateToken(username,password string) (string,error)
	ParseToken(token string) (int ,error)
}

type ToDoList interface {
	Create(userId int,list todoapp.ToDoList) (int,error)
	GetAll(userId int)([]todoapp.ToDoList,error)
	GetById(userId,listId int) (todoapp.ToDoList,error)
	Delete(userId,listId int) error
	Update(userId int,listId int,input todoapp.UpdateListInput) error
}

type TodoItem interface {
	Create(userId int,listId int,item todoapp.TodoItem) (int,error)
	GetAll(userId,listId int) ([]todoapp.TodoItem,error)
	GetById(userId,itemId int) (todoapp.TodoItem,error)
	Delete(userId,itemId int) error
	Update(userId,itemId int,input todoapp.UpdateItemInput) error
}

type Service struct {
	Authorization
	ToDoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ToDoList: NewTodoListService(repos.ToDoList),
		TodoItem: NewTodoItemService(repos.TodoItem,repos.ToDoList),
	}
}