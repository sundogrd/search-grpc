package search_test

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	searchGen "github.com/sundogrd/search-grpc/grpc_gen/search"
	"github.com/sundogrd/search-grpc/servers/search"
	"testing"
)

func newTestSearchServer_deleteIndex() (*search.SearchServiceServer, error) {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	return &search.SearchServiceServer{
		ElasticsearchClient: client,
	}, nil
}



func TestSearchServer_DeleteIndex(t *testing.T) {
	server, err := newTestSearchServer_deleteIndex()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("测试开始运行.....")
	res, err := server.AddIndex(context.Background(), &searchGen.AddIndexRequest{
		AppId: "lwio",
		IndexName: "test",
		IndexJson: "{}",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("AddIndex: %+v", res)
}
