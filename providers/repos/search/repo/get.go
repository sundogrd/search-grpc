package repo

import (
	"context"
	"fmt"

	repo "github.com/sundogrd/search-grpc/providers/repos/search"
)

func (s searchRepo) Get(ctx context.Context, req *repo.GetRequest) (*repo.GetResponse, error) {

	var search repo.Search
	db := s.gormDB

	dbc := db.Where(repo.Search{
		ID: req.SearchId,
	}).First(&search)

	if dbc.Error != nil {
		fmt.Printf("[providers/search] Delete: db get error: %+v", dbc.Error)
		return nil, dbc.Error
	}

	res := &repo.GetResponse{
		Search: &search,
	}

	if res.Search.ID == 0 {
		return nil, nil
	}

	return res, nil
}
