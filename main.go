package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pbfiles "shop-service/pbfiles"
	service "shop-service/service"

	"google.golang.org/grpc"
);

func main() {
	serverAddress := fmt.Sprintf(
		"%s:%s",
		"localhost",
		"50055",
	)
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	log.Printf("Starting server on  : %d", serverAddress)

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	srv := &service.CatalogServiceServer{}
	pbfiles.RegisterCatalogServiceServer(s, srv)

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	log.Println("Server successfully started")
	
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	
	log.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	log.Println("Done.")

}


