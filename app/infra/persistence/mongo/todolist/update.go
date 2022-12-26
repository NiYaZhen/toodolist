package todolist

import (
	"context"
	"time"
	model "todolist/todolist/app/domain/model/todolist"
	db "todolist/todolist/app/infra/persistence/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) Update(ctx context.Context, data *model.ToDoList) error {

	filter := bson.M{
		"id": data.TaskId,
	}
	data.UpdateAt = time.Now()
	update := bson.M{
		"$set": data,
	}

	if _, err := db.Ctodolist.UpdateOne(ctx, filter, update); err != nil {

	}
	return nil
}
