package main

import (
	"context"
	"github.com/guido4f/grpc-pricing/gen/byhiras.pricing"
	"github.com/guido4f/grpc-pricing/server/config"
	"github.com/guido4f/grpc-pricing/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	ctx         context.Context
	postService byhiras_pricing.PricingServiceServer
)

func init() {
	postService = service.NewPricingService(ctx)
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	startGrpcServer(config)
}

func startGrpcServer(config config.Config) {

	postServer := service.NewPricingService(ctx)

	grpcServer := grpc.NewServer()

	// ? Register the Post gRPC service
	byhiras_pricing.RegisterPricingServiceServer(grpcServer, postServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatal("cannot create gapi server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create gapi server: ", err)
	}
}
