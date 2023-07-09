package fees

type FeeRate interface {
	calculateFee(marketValue float64) float64
}

type basisPointFee struct {
	BasisPointRate float64
}

type percentageFee struct {
	PercentageFee float64
}

func newPercentageFee(percentage float64) *percentageFee {
	return &percentageFee{percentage}
}

type highWaterMarkFee struct {
	LowerBound float64
	Rate       FeeRate
}

type ComparisonFundFee struct {
	ComparisonFundStart  float64
	ComparisonFundSClose float64
	HurdlePercentage     float64
	FundToPriceStart     float64
	FundToPriceClose     float64
	DifferentialRate     FeeRate
}

type highWaterMarkedFee struct {
	Tiers []highWaterMarkFee
}

func (r percentageFee) calculateFee(v float64) float64 {
	return v * (r.PercentageFee / 100)
}
func (r basisPointFee) calculateFee(v float64) float64 {
	return v * (r.BasisPointRate / 10_000)
}
