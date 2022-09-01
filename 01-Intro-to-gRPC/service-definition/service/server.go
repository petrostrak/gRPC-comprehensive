// ProductInfo implementation with Go
package main

import (
	"context"

	pb "github.com/petrostrak/gRPC-comprehensive/01-Intro-to-gRPC/service-definition/service/ecommerce"
)

// Add product remote method
func AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {

	// business logic
	pid := pb.ProductID{
		Value: in.Id,
	}

	return &pid, nil
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
