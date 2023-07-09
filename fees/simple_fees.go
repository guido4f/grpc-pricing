package fees

type SimpleFeePricingRequest struct {
	BaseMarketValue float64
	Rate            FeeRate
}

type SimpleFeePricingResponse struct {
	SimpleFeePricingRequest
	calculatedFee float64
}

func Calculate(rate FeeRate, marketValue float64) float64 {
	return rate.calculateFee(marketValue)
}
