package search

import (
	"context"
	repository "todolist/todolist/app/domain/repo/todolist"
	"todolist/todolist/app/usecase/todolist"
)

type searchUsecase struct {
	todolistRepo repository.ToDoListRepo
}

func NewSearchUsecase(todolistRepo repository.ToDoListRepo) SearchUsecase {
	return &searchUsecase{
		todolistRepo: todolistRepo,
	}
}

func (u *searchUsecase) Search(ctx context.Context, input *SearchInput) (*SearchOutput, error) {
	output := new(SearchOutput)

	searchTask, err := u.todolistRepo.Search(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	output.Task = todolist.ConvertTaskDto(searchTask)
	return output, nil

}

func (u *searchUsecase) SearchType(ctx context.Context, input *SearchTypeInput) (*SearchTypeOutput, error) {
	output := new(SearchTypeOutput)

	searchTypeTask, _ := u.todolistRepo.SearchType(ctx, input.TaskType)

	output.List = todolist.ConvertTaskDtos(searchTypeTask)

	return output, nil

}
