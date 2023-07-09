package fees

type TierFee struct {
	UpperBound float64
	Rate       FeeRate
}

func NewTierFee(upperBound float64, rate FeeRate) *TierFee {
	return &TierFee{upperBound, rate}
}

type TieredFees struct {
	Tiers []TierFee
}

func (s TieredFees) Len() int {
	return len(s.Tiers)
}

func (s TieredFees) Swap(i, j int) {
	s.Tiers[i], s.Tiers[j] = s.Tiers[j], s.Tiers[i]
}

func (s TieredFees) Less(i, j int) bool {
	return s.Tiers[i].UpperBound < s.Tiers[j].UpperBound
}
