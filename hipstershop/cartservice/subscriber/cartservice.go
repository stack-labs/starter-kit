package subscriber

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	pb "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

type CartService struct{}

func (e *CartService) Handle(ctx context.Context, msg *pb.Empty) error {
	log.Log("Handler Received message: ")
	return nil
}

func Handler(ctx context.Context, msg *pb.Empty) error {
	log.Log("Function Received message: ")
	return nil
}
