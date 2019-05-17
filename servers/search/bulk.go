package search

import (
	"context"
	"github.com/pkg/errors"

	"github.com/sundogrd/search-grpc/grpc_gen/search"
)

func (server SearchServiceServer) Bulk(ctx context.Context, req *search.BulkRequest) (*search.BulkResponse, error) {

	return nil, errors.New("not implemented")
}
