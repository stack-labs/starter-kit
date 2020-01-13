package handler

import (
	"context"

	pb "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

type ShippingService struct{}

func (*ShippingService) GetQuote(context.Context, *pb.GetQuoteRequest, *pb.GetQuoteResponse) error {
	return nil
}
func (*ShippingService) ShipOrder(context.Context, *pb.ShipOrderRequest, *pb.ShipOrderResponse) error {
	return nil
}
