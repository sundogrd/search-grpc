package search

import (
	"context"
	"github.com/pkg/errors"

	"github.com/sundogrd/search-grpc/grpc_gen/search"
)

func (server SearchServiceServer) IndexDocument(ctx context.Context, req *search.IndexDocumentRequest) (*search.IndexDocumentResponse, error) {

	return nil, errors.New("not implemented")
}
