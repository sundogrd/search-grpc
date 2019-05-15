package search_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/sundogrd/search-grpc/grpc_gen/search"
	searchGen "github.com/sundogrd/search-grpc/grpc_gen/search"
)

// var message = make(chan bool)

func TestSearchServer_Like(t *testing.T) {

	// go initServer(message)
	// <-message
	fmt.Println("客户端开始运行.....")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	defer conn.Close()
	client := searchGen.NewSearchServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Like(ctx, &search.LikeRequest{
		SearchId: 343191254370103296,
	})

	log.Printf("%s: %s", name, res)

	if err != nil {
		t.Fatalf("Like Client err: %+v", err)
	}
	t.Logf("Like Client: %+v", res)
}
