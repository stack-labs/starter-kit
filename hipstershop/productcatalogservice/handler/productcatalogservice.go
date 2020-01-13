package handler

import (
	"context"

	pb "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

type ProductcatalogService struct{}

func (e *ProductcatalogService) ListProducts(context.Context, *pb.Empty, *pb.ListProductsResponse) error {

	return nil
}
func (e *ProductcatalogService) GetProduct(context.Context, *pb.GetProductRequest, *pb.Product) error {

	return nil
}
func (e *ProductcatalogService) SearchProducts(context.Context, *pb.SearchProductsRequest, *pb.SearchProductsResponse) error {

	return nil
}
