package service

import (
	todoapp "todo-app"
	"todo-app/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
	listRepo repository.ToDoList
}

func NewTodoItemService(repo repository.TodoItem,listRepo repository.ToDoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int,item todoapp.TodoItem) (int,error){
	_ , err := s.listRepo.GetById(userId,listId)
	if err != nil{
		//list does not exists or does not belongs to user
		return 0,err
	}

	return s.repo.Create(listId,item)

}

func (s *TodoItemService) GetAll(userId, listId int) ([]todoapp.TodoItem,error) {
	return s.repo.GetAll(userId,listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (todoapp.TodoItem,error){
	return s.repo.GetById(userId,itemId)
}

func(s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId,itemId)
}

func (s *TodoItemService) Update(userId,itemId int,input todoapp.UpdateItemInput) error {
	return s.repo.Update(userId,itemId,input)
}
