package get

import "todolist/todolist/app/usecase/todolist"

type GetInput struct {
	Useremail string
}

type GetOutput struct {
	List []*todolist.TaskDto
}
