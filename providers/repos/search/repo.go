package search

import (
	"context"
)

type GetRequest struct {
}
type GetResponse struct {
}

type Repo interface {
	Get(ctx context.Context, req *GetRequest) (*GetResponse, error)
}
