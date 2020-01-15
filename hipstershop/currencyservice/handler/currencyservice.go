package handler

import (
	"context"

	pb "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

type CurrencyService struct{}

func (e *CurrencyService) GetSupportedCurrencies(context.Context, *pb.Empty, *pb.GetSupportedCurrenciesResponse) error {

	return nil
}
func (e *CurrencyService) Convert(context.Context, *pb.CurrencyConversionRequest, *pb.Money) error {
	return nil
}
