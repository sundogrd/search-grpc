package search

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/sundogrd/search-grpc/grpc_gen/search"
)

func (server SearchServiceServer) AddIndex(ctx context.Context, req *search.AddIndexRequest) (*search.AddIndexResponse, error) {
	indicesCreateRes, err := server.ElasticsearchClient.Indices.Create(req.AppId + "_" + req.IndexName)
	if err != nil {
		logrus.Errorf("[search-grpc/servers/search] AddIndex err: %s", err.Error())
		return nil, err
	}
	if indicesCreateRes.StatusCode != 200 && indicesCreateRes.StatusCode != 204 {
		return nil, errors.New(fmt.Sprintf("[search-grpc/servers/search] AddIndex elastic response error: %s", indicesCreateRes.String()))
	}
	return &search.AddIndexResponse{
	}, nil
}
