package service

import (
	"context"
	"fmt"

	search "github.com/sundogrd/search-grpc/providers/repos/search"
	service "github.com/sundogrd/search-grpc/services/search"
)

func (s *searchService) Like(ctx context.Context, req *service.LikeRequest) (*service.LikeResponse, error) {
	repo := *s.searchRepo

	cmt, err := repo.Get(ctx, &search.GetRequest{
		SearchId: req.SearchId,
	})

	if err != nil {
		fmt.Printf("[service/search] Like: get error before update like count: %+v", err)
		return nil, err
	}

	response, err := repo.Update(ctx, &search.UpdateRequest{
		SearchId: req.SearchId,
		Map: map[string]interface{}{
			"like": cmt.Search.Like + 1,
		},
	})
	if err != nil {
		fmt.Printf("[service/search] Like: like error: %+v", err)
		return nil, err
	}
	res := &service.LikeResponse{
		SearchId: response.Search.ID,
	}
	return res, nil
}
