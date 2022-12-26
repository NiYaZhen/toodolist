package delete

import (
	"context"

	repository "todolist/todolist/app/domain/repo/todolist"
)

type deleteUsecase struct {
	todolistRepo repository.ToDoListRepo
}

func NewDeleteUsecase(todolistRepo repository.ToDoListRepo) DeleteUsecase {
	return &deleteUsecase{
		todolistRepo: todolistRepo,
	}
}

func (u *deleteUsecase) Delete(ctx context.Context, input *DeleteInput) error {

	return u.todolistRepo.Delete(ctx, input.ID)
}
