package update

import "context"

type UpdateUsecase interface {
	Update(ctx context.Context, input *UpdateInput) (*UpdateOutput, error)
}
