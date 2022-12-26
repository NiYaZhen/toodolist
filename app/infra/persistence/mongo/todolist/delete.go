package todolist

import (
	"context"
	db "todolist/todolist/app/infra/persistence/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) Delete(ctx context.Context, tasksID string) error {

	filter := bson.M{
		"id": tasksID,
	}

	if _, err := db.Ctodolist.DeleteOne(ctx, filter); err != nil {

	}
	return nil

}
