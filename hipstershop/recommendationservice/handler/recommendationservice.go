package handler

import (
	"context"

	pb "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

type RecommendationService struct{}

func (*RecommendationService) ListRecommendations(context.Context, *pb.ListRecommendationsRequest, *pb.ListRecommendationsResponse) error {
	return nil
}
