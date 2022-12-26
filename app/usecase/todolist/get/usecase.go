package get

import (
	"context"
	repository "todolist/todolist/app/domain/repo/todolist"
	service "todolist/todolist/app/domain/service/todolist"
	"todolist/todolist/app/usecase/todolist"
)

type getUsecase struct {
	todolistService service.TodoListService
	todolistRepo    repository.ToDoListRepo
}

func NewGetUsecase(todolistRepo repository.ToDoListRepo, todolistService service.TodoListService) GetUsecase {
	return &getUsecase{
		todolistRepo:    todolistRepo,
		todolistService: todolistService,
	}
}

func (u *getUsecase) Get(ctx context.Context, input *GetInput) (*GetOutput, error) {
	output := new(GetOutput)

	list, err := u.todolistRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	output.List = todolist.ConvertTaskDtos(list)
	return output, nil
}
