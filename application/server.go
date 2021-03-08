package grpc

import (
	pb "ayla-grpc/application/grpc/pb/protofiles"
	service "ayla-grpc/application/service"
	"ayla-grpc/model"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var productList = model.NewProducts()

func StartGrcpServer() {

	list, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("could not connect", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	productService := service.NewProductGrpcServer(productList)

	pb.RegisterProductServiceServer(grpcServer, productService)

	log.Println("gRPC server has been started!")

	if err := grpcServer.Serve(list); err != nil {
		log.Fatal("could not connect", err)
	}
}
