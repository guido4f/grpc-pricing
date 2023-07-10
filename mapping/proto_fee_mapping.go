package mapping

import (
	"PricingService/fees"
	byhiras_pricing "PricingService/gen/byhiras.pricing"
	"fmt"
)

func FromProtoTieredRate(proto *byhiras_pricing.PricingRequest_TieredRate) (fees.TieredFees, error) {

	tiers, err := FromProtoTiers(proto.TieredRate.Tiers)
	if err != nil {
		return fees.TieredFees{}, err
	}
	tieredFees := fees.TieredFees{
		tiers,
	}
	return tieredFees, nil
}
func FromProtoTiers(tiers []*byhiras_pricing.Tier) ([]fees.TierFee, error) {
	returnTiers := []fees.TierFee{}
	for _, tier := range tiers {
		protoTier, err := FromProtoTier(tier)
		if err != nil {
			return nil, err
		}
		returnTiers = append(returnTiers, protoTier)
	}
	return returnTiers, nil
}

func FromProtoTier(proto *byhiras_pricing.Tier) (fees.TierFee, error) {
	switch rate := proto.GetRate().Rate.(type) {
	case *byhiras_pricing.Rate_Basis:
		tier := fees.TierFee{
			*proto.HighBound,
			fees.NewBasisPointFee(rate.Basis),
		}
		return tier, nil
	case *byhiras_pricing.Rate_Percentage:
		tier := fees.TierFee{
			*proto.HighBound,
			fees.NewPercentageFee(rate.Percentage),
		}
		return tier, nil
	default:
		return fees.TierFee{}, fmt.Errorf("fromProtoTier: Type is not supported %s", rate)
	}
}
