package repo_test

import (
	"time"

	repo "github.com/sundogrd/search-grpc/providers/repos/search"
	searchRepo "github.com/sundogrd/search-grpc/providers/repos/search/repo"
	"github.com/sundogrd/gopkg/db"
)

func initTestDB() (repo.Repo, error) {

	gormDB, err := db.Connect(db.ConnectOptions{
		User:           "root",
		Password:       "12345678",
		Host:           "127.0.0.1",
		Port:           "3306",
		DBName:         "search",
		ConnectTimeout: "10s",
	})
	if err != nil {
		return nil, err
	}
	search, error := searchRepo.NewSearchRepo(gormDB, 2*time.Second)
	if error != nil {
		return nil, error
	}
	return search, nil
}
