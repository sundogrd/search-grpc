package search

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/sundogrd/search-grpc/grpc_gen/search"
)

func (server SearchServiceServer) IndexDocument(ctx context.Context, req *search.IndexDocumentRequest) (*search.IndexDocumentResponse, error) {
	createRes, err := server.ElasticsearchClient.Create(req.AppId + "_" + req.IndexName, strconv.FormatInt(req.DocId, 10), strings.NewReader(req.DocJson))
	if err != nil {
		logrus.Errorf("[search-grpc/servers/search] IndexDocument err: %s", err.Error())
		return nil, err
	}
	defer createRes.Body.Close()
	if createRes.IsError() {
		log.Printf("[%s] Error indexing document", createRes.Status())
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(createRes.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", createRes.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	if createRes.StatusCode != http.StatusCreated {
		return nil, errors.New(fmt.Sprintf("[search-grpc/servers/search] IndexDocument elastic response error: %s", createRes.String()))
	}
	return &search.IndexDocumentResponse{
		DocId: req.DocId,
	}, nil
}
