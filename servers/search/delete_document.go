package search

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strconv"

	"github.com/sundogrd/search-grpc/grpc_gen/search"
)

func (server SearchServiceServer) DeleteDocument(ctx context.Context, req *search.DeleteDocumentRequest) (*search.DeleteDocumentResponse, error) {
	deleteRes, err := server.ElasticsearchClient.Delete(req.AppId + "_" + req.IndexName, strconv.FormatInt(req.DocId, 10))
	if err != nil {
		logrus.Errorf("[search-grpc/servers/search] DeleteIndex err: %s", err.Error())
		return nil, err
	}
	defer deleteRes.Body.Close()
	if deleteRes.StatusCode != 200 && deleteRes.StatusCode != 204 {
		return nil, errors.New(fmt.Sprintf("[search-grpc/servers/search] DeleteDocument elastic response error: %s", deleteRes.String()))
	}
	return &search.DeleteDocumentResponse{
	}, nil
}
