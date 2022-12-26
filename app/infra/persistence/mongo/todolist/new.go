package todolist

import (
	repository "todolist/todolist/app/domain/repo/todolist"
)

type repo struct {
}

func NewRepo() repository.ToDoListRepo {
	return &repo{}
}
