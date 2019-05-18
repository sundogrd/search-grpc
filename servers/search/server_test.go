package search_test

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/sundogrd/search-grpc/servers/search"
)

func newTestSearchServer() (*search.SearchServiceServer, error) {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	return &search.SearchServiceServer{
		ElasticsearchClient: client,
	}, nil
}
