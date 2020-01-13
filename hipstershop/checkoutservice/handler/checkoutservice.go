package handler

import (
	"context"

	pb "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

type CheckoutService struct{}

func (e *CheckoutService) PlaceOrder(context.Context, *pb.PlaceOrderRequest, *pb.PlaceOrderResponse) error {
	return nil
}
