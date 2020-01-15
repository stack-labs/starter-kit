package handler

import (
	"context"

	pb "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

type AdService struct{}

func (e *AdService) GetAds(context.Context, *pb.AdRequest, *pb.AdResponse) error {
	return nil
}
