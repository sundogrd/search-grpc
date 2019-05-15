package repo_test

import (
	"context"
	"testing"

	repo "github.com/sundogrd/search-grpc/providers/repos/search"
)

func TestSearchProvider_Get(t *testing.T) {
	search, err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}

	res, err := search.Get(context.Background(), &repo.GetRequest{
		SearchId: 343087411107999744,
	})
	if err != nil {
		t.Fatalf("GetSearch err: %+v", err)
	}
	t.Logf("GetSearch: %+v", res)
}
