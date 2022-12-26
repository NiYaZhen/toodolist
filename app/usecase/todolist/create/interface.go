package create

import "context"

type CreateUsecase interface {
	Create(ctx context.Context, input *CreateInput) (*CreateOutput, error)
}
