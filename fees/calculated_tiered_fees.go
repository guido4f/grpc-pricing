package fees

type CalculatedTierFee struct {
	TierFee
	amountInTier float64
	fee          float64
}

func NewCalculatedTierFee(tierFee TierFee, amountInTier float64, fee float64) *CalculatedTierFee {
	return &CalculatedTierFee{TierFee: tierFee, amountInTier: amountInTier, fee: fee}
}

type CalculatedTieredFees struct {
	Tiers    []CalculatedTierFee
	totalFee float64
}

func (f *CalculatedTieredFees) update(tierFee TierFee, amountInTier float64) {
	calculatedFee := tierFee.Rate.calculateFee(amountInTier)

	fee := CalculatedTierFee{
		tierFee,
		amountInTier,
		calculatedFee,
	}

	f.Tiers = append(f.Tiers, fee)
	f.totalFee += calculatedFee
}
