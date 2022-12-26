package create

import (
	"context"
	"fmt"
	repository "todolist/todolist/app/domain/repo/todolist"
	service "todolist/todolist/app/domain/service/todolist"
)

type createUsecase struct {
	todolistService service.TodoListService
	todolistRepo    repository.ToDoListRepo
}

func NewCreateUsecase(todolistRepo repository.ToDoListRepo, todolistService service.TodoListService) CreateUsecase {
	return &createUsecase{
		todolistService: todolistService,
		todolistRepo:    todolistRepo,
	}
}

func (u *createUsecase) Create(ctx context.Context, input *CreateInput) (*CreateOutput, error) {
	output := new(CreateOutput)

	newTaskID, err := u.todolistService.NewTaskID(ctx)
	if err != nil {
		return nil, err
	}
	task, err := u.todolistService.NewTask(ctx, newTaskID, input.Title, input.Content, input.TaskType, input.Remark, input.StartdateAt, input.EnddateAt)
	fmt.Println(task)
	if err != nil {
		return nil, err
	}
	if err := u.todolistRepo.Create(ctx, task); err != nil {
		return nil, err
	}
	output.TaskId = newTaskID
	output.Title = task.Title
	output.Content = task.Content
	output.TaskType = task.TaskType
	output.Remark = task.Remark
	output.StartdateAt = task.StartdateAt
	output.StartdateAt = task.EnddateAt

	return output, nil
}
