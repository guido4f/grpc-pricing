package fees

type SimpleFee struct {
	BaseMarketValue float64
	Rate            FeeRate
}

type CalculatedSimpleFee struct {
	SimpleFee
	calculatedFee float64
}

func Calculate(rate FeeRate, marketValue float64) float64 {
	return rate.calculateFee(marketValue)
}
