package todolist

import (
	"context"
	model "todolist/todolist/app/domain/model/todolist"
	"todolist/todolist/app/infra"
	model2 "todolist/todolist/app/infra/model"
	db "todolist/todolist/app/infra/persistence/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) Search(ctx context.Context, taskID string) (*model.ToDoList, error) {
	filter := bson.M{
		"id": taskID,
	}
	var task *model.ToDoList

	if err := db.Ctodolist.FindOne(ctx, filter).Decode(&task); err != nil {
		return nil, nil
	}

	return task, nil

}

func (r *repo) SearchType(ctx context.Context, taskType string) ([]*model.ToDoList, error) {
	filter := bson.M{
		"tasktype": taskType,
	}
	var results []*model2.ToDoList
	d, _ := db.Ctodolist.Find(ctx, filter)
	if err := d.All(ctx, &results); err != nil {
		return nil, nil

	}
	return infra.ConvertToDomains(results), nil

}
