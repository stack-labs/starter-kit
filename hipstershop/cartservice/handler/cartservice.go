package handler

import (
	"context"

	pb "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

type CartService struct{}

func (e *CartService) AddItem(context.Context, *pb.AddItemRequest, *pb.Empty) error {
	return nil
}
func (e *CartService) GetCart(context.Context, *pb.GetCartRequest, *pb.Cart) error {
	return nil
}
func (e *CartService) EmptyCart(context.Context, *pb.EmptyCartRequest, *pb.Empty) error {
	return nil
}
