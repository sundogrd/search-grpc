package search_test

import (
	"context"
	"encoding/json"
	"fmt"
	searchGen "github.com/sundogrd/search-grpc/grpc_gen/search"
	"testing"
)

func TestSearchServiceServer_IndexDocument(t *testing.T) {
	server, err := newTestSearchServer()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("测试开始运行.....")
	doc := map[string]interface{}{
		"user": "fuc2kfg",
	}
	docJsonBytes, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	res, err := server.IndexDocument(context.Background(), &searchGen.IndexDocumentRequest{
		AppId:     "lwio",
		IndexName: "test",
		DocId: 15265532,
		DocJson: string(docJsonBytes),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("IndexDocument: %+v", res)
}
