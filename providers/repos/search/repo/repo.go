package repo

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sundogrd/search-grpc/providers/repos/search"
	"github.com/zheng-ji/goSnowFlake"
)

type searchRepo struct {
	gormDB         *gorm.DB
	contextTimeout time.Duration
	idBuilder      *goSnowFlake.IdWorker
}

// NewUserService will create new an articleUsecase object representation of article.Usecase interface
func NewSearchRepo(gormDB *gorm.DB, timeout time.Duration) (search.Repo, error) {
	idBuilder, err := goSnowFlake.NewIdWorker(3)
	if err != nil {
		return nil, err
	}

	hasTable := gormDB.HasTable(&search.Search{})
	if hasTable == false {
		gormDB.CreateTable(&search.Search{})
	} else {
		gormDB.AutoMigrate(&search.Search{})
	}

	repo := searchRepo{
		gormDB:         gormDB,
		contextTimeout: timeout,
		idBuilder:      idBuilder,
	}
	return repo, nil
}
