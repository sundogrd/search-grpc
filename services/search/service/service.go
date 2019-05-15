package service

import (
	"time"

	searchRepo "github.com/sundogrd/search-grpc/providers/repos/search"
	"github.com/sundogrd/search-grpc/services/search"
)

type searchService struct {
	searchRepo    *searchRepo.Repo
	contextTimeout time.Duration
}

// NewSearchService will create new an articleUsecase object representation of article.Usecase interface
func NewSearchService(searchRepo *searchRepo.Repo, timeout time.Duration) (search.Service, error) {
	return &searchService{
		searchRepo:    searchRepo,
		contextTimeout: timeout,
	}, nil
}
