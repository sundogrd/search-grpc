package main

import (
	"github.com/sirupsen/logrus"
	"net"
	"time"

	searchGen "github.com/sundogrd/search-grpc/grpc_gen/search"
	searchRepo "github.com/sundogrd/search-grpc/providers/repos/search/repo"
	"github.com/sundogrd/search-grpc/servers/search"
	searchService "github.com/sundogrd/search-grpc/services/search/service"
	configUtils "github.com/sundogrd/gopkg/config"
	"github.com/sundogrd/gopkg/db"
	grpcUtils "github.com/sundogrd/gopkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := configUtils.ReadConfigFromFile("./config", nil)
	if err != nil {
		logrus.Errorf("[search-grpc] ReadConfigFromFile err: %s", err.Error())
		panic(err)
	}

	instanceAddr := config.Get("grpcService.host").(string) + ":" + config.Get("grpcService.port").(string)
	listen, err := net.Listen("tcp", instanceAddr)
	if err != nil {
		logrus.Errorf("[search-grpc] net.Listen err: %s", err.Error())
		panic(err)
	}

	gormDB, err := db.Connect(db.ConnectOptions{
		User:           config.Get("db.options.user").(string),
		Password:       config.Get("db.options.password").(string),
		Host:           config.Get("db.options.host").(string),
		Port:           config.Get("db.options.port").(string),
		DBName:         config.Get("db.options.dbname").(string),
		ConnectTimeout: config.Get("db.options.connectTimeout").(string),
	})
	if err != nil {
		logrus.Errorf("[search-grpc] db.Connect err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[search-grpc] db.Connect finished")

	cr, err := searchRepo.NewSearchRepo(gormDB, 2*time.Second)
	if err != nil {
		logrus.Errorf("[search-grpc] NewSearchRepo err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[search-grpc] NewSearchRepo finished")

	cs, err := searchService.NewSearchService(&cr, 2*time.Second)
	if err != nil {
		logrus.Errorf("[search-grpc] NewSearchService err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[search-grpc] NewSearchService finished")

	grpcServer := grpc.NewServer()
	resolver, err := grpcUtils.NewGrpcResolover()
	if err != nil {
		logrus.Errorf("[search-grpc] NewGrpcResolover err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[search-grpc] NewGrpcResolover finished")

	err = grpcUtils.ResgiterServer(*resolver, "sundog.search", instanceAddr, 5*time.Second, 5)
	if err != nil {
		logrus.Errorf("[search-grpc] RegisterServer err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[search-grpc] ResgiterServer finished, service: %s, %s", "sundog.search", instanceAddr)

	searchGen.RegisterSearchServiceServer(grpcServer, &search.SearchServiceServer{
		GormDB:         gormDB,
		SearchRepo:    cr,
		SearchService: cs,
	})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}
