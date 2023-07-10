package main

import (
	proto "FeeCalculatorService/gen/guido4f.fee"
	"FeeCalculatorService/server/config"
	"FeeCalculatorService/server/service"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	ctx         context.Context
	postService proto.FeeCalculatorServiceServer
)

func init() {
	postService = service.NewPricingService(ctx)
}

func main() {
	cfg, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load cfg", err)
	}

	startGrpcServer(cfg)
}

func startGrpcServer(appCfg config.Config) {

	postServer := service.NewPricingService(ctx)

	grpcServer := grpc.NewServer()

	// ? Register the Post gRPC service
	proto.RegisterFeeCalculatorServiceServer(grpcServer, postServer)
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
