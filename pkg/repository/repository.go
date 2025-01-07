package repository

import (
	todoapp "todo-app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todoapp.User) (int,error)
	GetUser(username, password string) (todoapp.User, error)
}

type ToDoList interface {
	Create(userId int, list todoapp.ToDoList) (int,error)
	GetAll(userId int) ([]todoapp.ToDoList,error)
	GetById(userId, listId int) (todoapp.ToDoList,error)
	Delete(userId,listId int) error
	Update(userId,listId int,input todoapp.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int,item todoapp.TodoItem) (int,error)
	GetAll(userId,listId int) ([]todoapp.TodoItem,error)
	GetById(userId, itemId int) (todoapp.TodoItem,error)
	Delete(userId,itemId int) error
	Update(userId,listId int,input todoapp.UpdateItemInput) error
}

type Repository struct {
	Authorization
	ToDoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ToDoList: NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}