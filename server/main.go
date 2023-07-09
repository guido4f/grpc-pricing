package main

import (
	"PricingService/gen/byhiras.pricing"
	"PricingService/server/config"
	"PricingService/server/service"
	"context"
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

func startGrpcServer(appCfg config.Config) {

	postServer := service.NewPricingService(ctx)

	grpcServer := grpc.NewServer()

	// ? Register the Post gRPC service
	byhiras_pricing.RegisterPricingServiceServer(grpcServer, postServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", appCfg.Grpc.Host+":"+appCfg.Grpc.Port)
	if err != nil {
		log.Fatal("cannot create gapi server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create gapi server: ", err)
	}
}
