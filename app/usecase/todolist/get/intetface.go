package get

import "context"

type GetUsecase interface {
	Get(ctx context.Context, input *GetInput) (*GetOutput, error)
}
