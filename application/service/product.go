package service

import (
	pb "ayla-grpc/application/grpc/pb/protofiles"
	"ayla-grpc/model"
	"context"
	"time"
)

type ProductGrpcServer struct {
	pb.UnimplementedProductServiceServer
	Products *model.Products
}

func (p *ProductGrpcServer) CreateProduct(ctx context.Context, in *pb.Product) (*pb.ProductResult, error) {
	product := model.NewProduct()
	product.Name = in.Name
	p.Products.Add(product)

	return &pb.ProductResult{
		Id:   product.ID,
		Name: product.Name,
	}, nil
}

func (p *ProductGrpcServer) List(in *pb.Empty, stream pb.ProductService_ListServer) error {

	for _, product := range p.Products.Product {
		time.Sleep(time.Second * 5)
		stream.Send(&pb.ProductResult{Id: product.ID, Name: product.Name})
	}
	return nil
}

func NewProductGrpcServer(products *model.Products) *ProductGrpcServer {
	return &ProductGrpcServer{
		Products: products,
	}
}
