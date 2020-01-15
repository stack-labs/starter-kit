package handler

import (
	"context"

	pb "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

type PaymentService struct{}

func (*PaymentService) Charge(context.Context, *pb.ChargeRequest, *pb.ChargeResponse) error {
	return nil
}
