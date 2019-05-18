package search_test

import (
	"context"
	"encoding/json"
	"fmt"
	searchGen "github.com/sundogrd/search-grpc/grpc_gen/search"
	"testing"
)

func TestSearchServiceServer_Search(t *testing.T) {
	server, err := newTestSearchServer()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("测试开始运行.....")
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"term" : map[string]interface{}{
				"user" : map[string]interface{}{
					"value": "fuc2kfg",
					"boost": 1.0,
				},
			},
		},
	}
	queryJsonBytes, err := json.Marshal(query)
	if err != nil {
		t.Fatal(err)
	}
	res, err := server.Search(context.Background(), &searchGen.SearchRequest{
		AppId:     "lwio",
		IndexName: "test",
		QueryJson: string(queryJsonBytes),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Search: %+v", res)
}