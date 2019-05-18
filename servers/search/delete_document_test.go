package search_test

import (
	"context"
	"fmt"
	searchGen "github.com/sundogrd/search-grpc/grpc_gen/search"
	"testing"
)

func TestSearchServer_DeleteDocument(t *testing.T) {
	server, err := newTestSearchServer()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("测试开始运行.....")
	res, err := server.DeleteDocument(context.Background(), &searchGen.DeleteDocumentRequest{
		AppId:     "lwio",
		IndexName: "test",
		DocId: 443474713,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("DeleteDocument: %+v", res)
}
