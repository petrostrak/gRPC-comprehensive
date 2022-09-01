// ProductInfo implementation with Go
package main

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	pb "github.com/petrostrak/gRPC-comprehensive/01-Intro-to-gRPC/service-definition/service/ecommerce"
)

// server is used to implement ecommerce/product_info.
type server struct {
	productMap map[string]*pb.Product
}

// Add product remote method
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("error while generating Product ID")
	}

	in.Id = id.String()

	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in

	return &pb.ProductID{Value: in.Id}, nil
}

// Get product remote method
func getProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	// business logic
	p := pb.Product{
		Id:          in.Value,
		Name:        "prod",
		Description: "This is a product",
	}

	return &p, nil
}
