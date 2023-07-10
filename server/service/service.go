package service

import (
	fees "FeeCalculatorService/gen/guido4f.fee"
	"FeeCalculatorService/mapping"
	"context"
	"fmt"
)

type CalculateFeeServiceImpl struct {
	ctx context.Context
}

func NewPricingService(ctx context.Context) fees.FeeCalculatorServiceServer {
	return &CalculateFeeServiceImpl{ctx}
}

func (p CalculateFeeServiceImpl) CalculateFee(ctx context.Context, request *fees.FeeCalculatorRequest) (*fees.FeeCalculatorResponse, error) {
	switch t := request.GetRates().(type) {
	case *fees.FeeCalculatorRequest_PerformanceFee:
		fmt.Printf("PerformanceFDees")
	case *fees.FeeCalculatorRequest_ScaledRate:

	case *fees.FeeCalculatorRequest_TieredRate:
		_, err := mapping.FromProtoTieredRate(t)
		if err != nil {
			//TODO Response with an Error
		}
		//TODO Respond with the Rate

	default:
		fmt.Printf("No matching operations %s", t)
	}
	return nil, nil
}
