// ProductInfo implementation with Go
package gRPC_comprehensive

import (
	"context"
)

// Add product remote method
func AddProduct(ctx context.Context, in *Product) (*ProductID, error) {

	// business logic
	pid := ProductID{
		Value: in.Id,
	}

	return &pid, nil
}

// Get product remote method
func getProduct(ctx context.Context, in *ProductID) (*Product, error) {
	// business logic
	p := Product{
		Id:          in.Value,
		Name:        "prod",
		Description: "This is a product",
	}

	return &p, nil
}
