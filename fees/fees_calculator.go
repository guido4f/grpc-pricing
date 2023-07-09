package fees

import "sort"

func (s TieredFees) calculateTierFee(amount float64) CalculatedTieredFees {
	sort.Sort(&s)
	var lowerBound float64 = 0

	fees := CalculatedTieredFees{
		[]CalculatedTierFee{},
		0.0,
	}

	for _, t := range s.Tiers {

		if t.UpperBound < amount {
			fees.update(t, t.UpperBound-lowerBound)
			lowerBound = t.UpperBound
		} else {
			fees.update(t, amount-lowerBound)
			return fees
		}
	}
	return fees

}
