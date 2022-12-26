package todolist

import (
	"time"
	model "todolist/todolist/app/domain/model/todolist"
)

type TaskDto struct {
	Title       string
	Content     string
	Remark      string
	TaskType    string
	StartdateAt string
	EnddateAt   string
}

func ConvertTaskDto(in *model.ToDoList) *TaskDto {
	if in == nil {
		return new(TaskDto)
	}

	return &TaskDto{
		Title:       in.Title,
		Content:     in.Content,
		Remark:      in.TaskType,
		StartdateAt: in.StartdateAt.Format(time.RFC3339),
		EnddateAt:   in.EnddateAt.Format(time.RFC3339),
	}
}

func ConvertTaskDtos(in []*model.ToDoList) []*TaskDto {
	ans := make([]*TaskDto, len(in))

	for i, v := range in {
		ans[i] = ConvertTaskDto(v)
	}
	return ans

}
