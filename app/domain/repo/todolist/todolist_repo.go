package repo

import (
	"context"
	model "todolist/todolist/app/domain/model/todolist"
)

type ToDoListRepo interface {
	Create(ctx context.Context, todolist *model.ToDoList) error
	Update(ctx context.Context, todolist *model.ToDoList) error
	Get(ctx context.Context) ([]*model.ToDoList, error)
	Delete(ctx context.Context, tasksID string) error
	Search(ctx context.Context, tasksID string) (*model.ToDoList, error)
	SearchType(ctx context.Context, taskType string) ([]*model.ToDoList, error)
}
