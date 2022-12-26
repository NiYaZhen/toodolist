package service

import (
	"context"
	model "todolist/todolist/app/domain/model/todolist"
)

type TodoListService interface {
	NewTaskID(ctx context.Context) (string, error)
	NewTask(ctx context.Context, taskId, titie, content, tasktype, remark, startdateAt, enddateAt string) (*model.ToDoList, error)
	NewUpdateTask(ctx context.Context, titie, content, tasktype, remark, startdateAt, enddateAt string) (*model.ToDoList, error)
	GetTaskStatus(ctx context.Context, startdateAt, enddateAt string) (int32, error)
	GetTimeleft(ctx context.Context, enddateAt string) (string, error)
}
