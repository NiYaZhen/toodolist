package update

import (
	"context"
	repository "todolist/todolist/app/domain/repo/todolist"
	service "todolist/todolist/app/domain/service/todolist"
)

type updateUsecase struct {
	todolistService service.TodoListService
	todolistRepo    repository.ToDoListRepo
}

func NewUpdateUsecase(todolistRepo repository.ToDoListRepo, todolistService service.TodoListService) UpdateUsecase {
	return &updateUsecase{
		todolistService: todolistService,
		todolistRepo:    todolistRepo,
	}
}

func (u *updateUsecase) Update(ctx context.Context, input *UpdateInput) (*UpdateOutput, error) {
	task, err := u.todolistRepo.Search(ctx, input.ID)
	output := new(UpdateOutput)
	if err != nil {
		return nil, err
	}

	updatetask, err := u.todolistService.NewUpdateTask(ctx, input.Title, input.Content, input.TaskType, input.Remark, input.StartdateAt, input.EnddateAt)
	task.Title = updatetask.Title
	task.Content = updatetask.Content
	task.TaskType = updatetask.TaskType
	task.Remark = updatetask.Remark
	task.StartdateAt = updatetask.StartdateAt
	task.EnddateAt = updatetask.EnddateAt

	task.Status, err = u.todolistService.GetTaskStatus(ctx, input.StartdateAt, input.EnddateAt)
	task.Timeleft, err = u.todolistService.GetTimeleft(ctx, input.EnddateAt)

	output.Title = updatetask.Title
	output.Content = updatetask.Content
	output.TaskType = updatetask.TaskType
	output.Remark = updatetask.Remark
	output.StartdateAt = updatetask.StartdateAt
	output.EnddateAt = updatetask.EnddateAt

	output.Status, err = u.todolistService.GetTaskStatus(ctx, input.StartdateAt, input.EnddateAt)
	output.Timeleft, err = u.todolistService.GetTimeleft(ctx, input.EnddateAt)

	if err := u.todolistRepo.Update(ctx, task); err != nil {
		return nil, err
	}

	return output, nil

}

func (u *updateUsecase) Complete(ctx context.Context, input *CompleteInput) {

}
