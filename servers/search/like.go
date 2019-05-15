package search

import (
	"context"
	"github.com/pkg/errors"

	"github.com/sundogrd/search-grpc/grpc_gen/search"
)

func (server SearchServiceServer) Like(ctx context.Context, req *search.LikeRequest) (*search.LikeResponse, error) {

	return nil, errors.New("not implemented")
}
