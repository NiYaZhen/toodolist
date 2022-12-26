package todolist

import (
	"context"
	model "todolist/todolist/app/domain/model/todolist"
	db "todolist/todolist/app/infra/persistence/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) Get(ctx context.Context) ([]*model.ToDoList, error) {

	cursor, _ := db.Ctodolist.Find(ctx, bson.D{{}})

	var results []*model.ToDoList
	if err := cursor.All(ctx, &results); err != nil {
		return nil, nil
	}

	return results, nil

}
