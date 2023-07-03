package main

import (
	"../config"
	"./service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	ctx context.Context

	// ? Create the Post Variables
	postService services.PatchService
)

func init() {
	//config, err := config.LoadConfig(".")
	//if err != nil {
	//	log.Printf("Could not load environment variables", err)
	//}

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

	postServer, err := gapi.NewGrpcPostServer(postCollection, postService)
	if err != nil {
		log.Fatal("cannot create gapi postServer: ", err)
	}
	grpcServer := grpc.NewServer()

	// ? Register the Post gRPC service
	pb.RegisterPatchServiceServer(grpcServer, postServer)
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
