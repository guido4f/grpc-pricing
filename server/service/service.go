package service

import (
	byhiras_pricing "PricingService/gen/byhiras.pricing"
	"context"
	"fmt"
)

type PricingServiceImpl struct {
	ctx context.Context
}

func NewPricingService(ctx context.Context) byhiras_pricing.PricingServiceServer {
	return &PricingServiceImpl{ctx}
}

func (p PricingServiceImpl) Patch(ctx context.Context, request *byhiras_pricing.PricingRequest) (*byhiras_pricing.PricingResponse, error) {
	switch op := request.GetRates().(type) {
	case *byhiras_pricing.PricingRequest_PerformanceFee:
		fmt.Printf("Copy Operation start: %d, end : %d\n", op.CopyOp.Start, op.CopyOp.End)
	case *byhiras_pricing.PricingRequest_TieredRate:

	case *byhiras_pricing.PricingRequest_ScaledRate:
	default:
		fmt.Println("No matching operations")
	}
}
