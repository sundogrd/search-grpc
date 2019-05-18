package search

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/jinzhu/gorm"
)

type SearchServiceServer struct {
	GormDB              *gorm.DB
	ElasticsearchClient *elasticsearch.Client
}
