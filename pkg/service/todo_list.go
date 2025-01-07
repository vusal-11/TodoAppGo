package service

import (
	todoapp "todo-app"
	"todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.ToDoList
}

func NewTodoListService(repo repository.ToDoList) *TodoListService {
	return &TodoListService{repo: repo}
}


func (s *TodoListService) Create(userId int, list todoapp.ToDoList) (int,error) {
	return s.repo.Create(userId,list)
}

func (s *TodoListService) GetAll(userId int) ([]todoapp.ToDoList,error) {
	return s.repo.GetAll(userId)
}

func(s *TodoListService) GetById(userId, listId int) (todoapp.ToDoList,error) {
	return s.repo.GetById(userId,listId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input todoapp.UpdateListInput ) error{
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId,listId,input)
}