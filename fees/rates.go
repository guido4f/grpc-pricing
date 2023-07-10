package fees

type FeeRate interface {
	calculateFee(marketValue float64) float64
}

type BasisPointFee struct {
	BasisPointRate float64
}

func NewBasisPointFee(basisPointRate float64) *BasisPointFee {
	return &BasisPointFee{BasisPointRate: basisPointRate}
}

type PercentageFee struct {
	PercentageFee float64
}

func NewPercentageFee(percentage float64) *PercentageFee {
	return &PercentageFee{percentage}
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

func (r PercentageFee) calculateFee(v float64) float64 {
	return v * (r.PercentageFee / 100)
}
func (r BasisPointFee) calculateFee(v float64) float64 {
	return v * (r.BasisPointRate / 10_000)
}
