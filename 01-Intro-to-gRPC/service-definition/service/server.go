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

// AddProduct implements ecommerce.AddProduct
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

// GetProduct implements ecommerce.GetProduct
func (s *server) getProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if !exists {
		return nil, fmt.Errorf("Product does not exist")
	}

	return value, nil
}
