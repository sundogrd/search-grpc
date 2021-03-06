package search_test

import (
	"context"
	"fmt"
	searchGen "github.com/sundogrd/search-grpc/grpc_gen/search"
	"testing"
)

func TestSearchServer_DeleteIndex(t *testing.T) {
	server, err := newTestSearchServer()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("测试开始运行.....")
	res, err := server.DeleteIndex(context.Background(), &searchGen.DeleteIndexRequest{
		AppId:     "lwio",
		IndexName: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("DeleteIndex: %+v", res)
}
