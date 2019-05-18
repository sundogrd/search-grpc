package search_test

import (
	"context"
	"fmt"
	searchGen "github.com/sundogrd/search-grpc/grpc_gen/search"
	"testing"
)

func TestSearchServer_AddIndex(t *testing.T) {
	server, err := newTestSearchServer()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("测试开始运行.....")
	res, err := server.AddIndex(context.Background(), &searchGen.AddIndexRequest{
		AppId:     "lwio",
		IndexName: "test",
		IndexJson: "{}",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("AddIndex: %+v", res)
}
