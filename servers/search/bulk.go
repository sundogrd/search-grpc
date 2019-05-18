package search

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sundogrd/search-grpc/grpc_gen/search"
)

func (server SearchServiceServer) Bulk(ctx context.Context, req *search.BulkRequest) (*search.BulkResponse, error) {
	//indicesCreateRes, err := server.ElasticsearchClient.Bulk(req.)
	//if err != nil {
	//	logrus.Errorf("[search-grpc/servers/search] AddIndex err: %s", err.Error())
	//	return nil, err
	//}
	//if indicesCreateRes.StatusCode != 200 && indicesCreateRes.StatusCode != 204 {
	//	return nil, errors.New(fmt.Sprintf("[search-grpc/servers/search] AddIndex elastic response error: %s", indicesCreateRes.String()))
	//}
	//return &search.AddIndexResponse{
	//}, nil
	return nil, errors.New("not implemented")
}
