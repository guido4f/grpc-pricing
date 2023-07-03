package service

import (
	byhiras_pricing "../../gen/byhiras.pricing"
	"context"
)

type PricingServiceImpl struct {
	ctx context.Context
}

func NewPricingService(ctx context.Context) byhiras_pricing.PricingServiceServer {
	return &PricingServiceImpl{ctx}
}

func (p PricingServiceImpl) Patch(ctx context.Context, request *byhiras_pricing.PricingRequest) (*byhiras_pricing.PricingResponse, error) {
	//TODO implement me
	panic("implement me")
}
