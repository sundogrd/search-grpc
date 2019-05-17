package search

import (
	"context"
	"github.com/pkg/errors"

	"github.com/sundogrd/search-grpc/grpc_gen/search"
)

func (server SearchServiceServer) Search(ctx context.Context, req *search.SearchRequest) (*search.SearchResponse, error) {
	return nil, errors.New("not implemented")
}