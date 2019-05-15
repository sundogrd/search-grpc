package search

import (
	"context"
)

type LikeRequest struct {
	SearchId int64
}

type LikeResponse struct {
	SearchId int64
}

type Service interface {
	Like(ctx context.Context, req *LikeRequest) (*LikeResponse, error)
}
