package mapping

import (
	"FeeCalculatorService/fees"
	proto "FeeCalculatorService/gen/guido4f.fee"
	"fmt"
)

func FromProtoTieredRate(proto *proto.FeeCalculatorRequest_TieredRate) (fees.TieredFees, error) {

	tiers, err := FromProtoTiers(proto.TieredRate.Tiers)
	if err != nil {
		return fees.TieredFees{}, err
	}
	tieredFees := fees.TieredFees{
		tiers,
	}
	return tieredFees, nil
}
func FromProtoTiers(tiers []*proto.Tier) ([]fees.TierFee, error) {
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

func FromProtoTier(protoTier *proto.Tier) (fees.TierFee, error) {
	switch rate := protoTier.GetRate().Rate.(type) {
	case *proto.Rate_Basis:
		tier := fees.TierFee{
			*protoTier.HighBound,
			fees.NewBasisPointFee(rate.Basis),
		}
		return tier, nil
	case *proto.Rate_Percentage:
		tier := fees.TierFee{
			*protoTier.HighBound,
			fees.NewPercentageFee(rate.Percentage),
		}
		return tier, nil
	default:
		return fees.TierFee{}, fmt.Errorf("fromProtoTier: Type is not supported %s", rate)
	}
}
