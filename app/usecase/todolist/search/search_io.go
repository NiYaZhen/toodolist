package search

import "todolist/todolist/app/usecase/todolist"

type SearchInput struct {
	ID string
}

type SearchOutput struct {
	Task *todolist.TaskDto
}
type SearchTypeInput struct {
	TaskType string
}

type SearchTypeOutput struct {
	List []*todolist.TaskDto
}
