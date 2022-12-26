package infra

import (
	model "todolist/todolist/app/domain/model/todolist"
	model2 "todolist/todolist/app/infra/model"
)

func ConvertToDomains(in []*model2.ToDoList) []*model.ToDoList {
	ans := make([]*model.ToDoList, len(in))
	for i, v := range in {
		ans[i] = ConvertToDomain(v)
	}
	return ans
}

func ConvertToDomain(in *model2.ToDoList) *model.ToDoList {
	return &model.ToDoList{
		TaskId:      in.TaskId,
		Title:       in.Title,
		Content:     in.Content,
		Remark:      in.Remark,
		TaskType:    in.TaskType,
		UserEmail:   in.UserEmail,
		Status:      in.Status,
		Complete:    in.Complete,
		CreatedAt:   in.CreatedAt,
		UpdateAt:    in.UpdateAt,
		StartdateAt: in.StartdateAt,
		EnddateAt:   in.EnddateAt,
		CompletedAt: in.CompletedAt,
		Timeleft:    in.Timeleft,
	}
}
