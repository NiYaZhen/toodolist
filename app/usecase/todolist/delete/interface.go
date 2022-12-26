package delete

import "context"

type DeleteUsecase interface {
	Delete(ctx context.Context, input *DeleteInput) error
}
