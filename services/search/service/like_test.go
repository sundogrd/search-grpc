package service_test

import (
	"context"
	"testing"

	searchService "github.com/sundogrd/search-grpc/services/search"
)

func TestSearchService_Like(t *testing.T) {

	cs, err := initServiceTest()
	if err != nil {
		t.Fatalf("Like Service err: %+v", err)
	}
	res, err := cs.Like(context.Background(), &searchService.LikeRequest{
		SearchId: 343181982320046080,
	})
	if err != nil {
		t.Fatalf("Like Service err: %+v", err)
	}
	t.Logf("Like Service: %+v", res)

}
