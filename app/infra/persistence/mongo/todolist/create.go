package todolist

import (
	"context"
	model "todolist/todolist/app/domain/model/todolist"
	db "todolist/todolist/app/infra/persistence/mongo"
)

func (r *repo) Create(ctx context.Context, data *model.ToDoList) error {

	if _, err := db.Ctodolist.InsertOne(ctx, data); err != nil {

	}

	return nil

}
