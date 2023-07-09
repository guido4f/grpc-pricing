package fees

import "sort"

type TierFee struct {
	UpperBound float64
	Rate       FeeRate
}

type TieredFees struct {
	Tiers []TierFee
}

type PricedTierFee struct {
	TierFee
	amountInTier float64
	fee          float64
}

type PricedTieredFees struct {
	Tiers    []PricedTierFee
	totalFee float64
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

func (s TieredFees) calculateTierFee(amount float64) PricedTieredFees {
	sort.Sort(&s)
	var lowerBound float64 = 0

	fees := PricedTieredFees{
		[]PricedTierFee{},
		0.0,
	}

	for _, t := range s.Tiers {

		if t.UpperBound < amount {
			var amountInTier float64 = t.UpperBound - lowerBound
			fees.update(t, amountInTier)
			lowerBound = t.UpperBound
		} else {
			var amountInTier float64 = amount - lowerBound
			fees.update(t, amountInTier)
			return fees
		}
	}
	return fees

}

func (calculatedFees PricedTieredFees) update(tierFee TierFee, amountInTier float64) {
	calculatedFee := tierFee.Rate.calculateFee(amountInTier)

	var fee = PricedTierFee{
		tierFee,
		amountInTier,
		calculatedFee,
	}

	calculatedFees.Tiers = append(calculatedFees.Tiers, fee)
	calculatedFees.totalFee += calculatedFee
}
