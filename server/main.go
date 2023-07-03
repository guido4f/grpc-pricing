package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Printf("Could not load environment variables", err)
	}

	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// ? Instantiate the Constructors
	postCollection = mongoclient.Database("golang_mongodb").Collection("posts")
	postService = services.NewPostService(postCollection, ctx)

}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer mongoclient.Disconnect(ctx)

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
