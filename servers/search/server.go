package search

import (
	"github.com/jinzhu/gorm"
	searchRepo "github.com/sundogrd/search-grpc/providers/repos/search"
	searchService "github.com/sundogrd/search-grpc/services/search"
	"github.com/elastic/go-elasticsearch/v7"
)

type SearchServiceServer struct{
	GormDB *gorm.DB
	SearchRepo searchRepo.Repo
	SearchService searchService.Service
	ElasticsearchClient *elasticsearch.Client
}