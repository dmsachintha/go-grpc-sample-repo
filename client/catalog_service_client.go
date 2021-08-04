package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pbfiles "shop-service/pbfiles"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := fmt.Sprintf(
		"%s:%s",
		"localhost",
		"50055",
	)
	log.Printf("dial server %s", serverAddress)

	connx, errx := grpc.Dial(serverAddress, grpc.WithInsecure())
	if errx != nil {
		log.Fatal("cannot dial server:", errx)
	}

	catalogClient := pbfiles.NewCatalogServiceClient(connx)

	ctxX, cancelX := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelX()

	catalogListRequest := &pbfiles.ViewCatalogRequest{
		CatalogType: "books",
	}

	result, errX := catalogClient.List(ctxX, catalogListRequest)

	if errX != nil {
		fmt.Println("error is ", errX)
	}
	log.Println(result)
}
