package search

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/sundogrd/search-grpc/grpc_gen/search"
)

func (server SearchServiceServer) DeleteIndex(ctx context.Context, req *search.DeleteIndexRequest) (*search.DeleteIndexResponse, error) {
	indicesDeleteRes, err := server.ElasticsearchClient.Indices.Delete([]string{req.AppId + "_" + req.IndexName})
	if err != nil {
		logrus.Errorf("[search-grpc/servers/search] DeleteIndex err: %s", err.Error())
		return nil, err
	}
	defer indicesDeleteRes.Body.Close()
	if indicesDeleteRes.StatusCode != 200 && indicesDeleteRes.StatusCode != 204 {
		return nil, errors.New(fmt.Sprintf("[search-grpc/servers/search] DeleteIndex elastic response error: %s", indicesDeleteRes.String()))
	}
	return &search.DeleteIndexResponse{
	}, nil
}
