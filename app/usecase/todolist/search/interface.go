package search

import "context"

type SearchUsecase interface {
	Search(ctx context.Context, input *SearchInput) (*SearchOutput, error)
	SearchType(ctx context.Context, input *SearchTypeInput) (*SearchTypeOutput, error)
}
