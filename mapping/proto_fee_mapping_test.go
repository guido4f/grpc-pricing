package mapping

import (
	"PricingService/fees"
	byhiras_pricing "PricingService/gen/byhiras.pricing"
	"testing"
)

func TestFromProtoTieredRate(t *testing.T) {
	type args struct {
		proto *byhiras_pricing.PricingRequest_TieredRate
	}
	feeRate := fees.NewPercentageFee(64.0)
	percentage := byhiras_pricing.Rate_Percentage{20}
	rate := byhiras_pricing.Rate{Rate: &percentage}
	bound := 10.0
	tier := byhiras_pricing.Tier{
		HighBound: &bound,
		Rate:      &rate,
	}

	tiers := &[]byhiras_pricing.Tier{tier,}

	request_tieredRate := byhiras_pricing.PricingRequest_TieredRate{
		TieredRate: {
			Tiers: tiers,
		},
	}
	requestTieredRate := request_tieredRate
	tieredRate := requestTieredRate
	rate := tieredRate
	tests := []struct {
		name    string
		args    args
		want    fees.TieredFees
		wantErr bool
	}{
		{name: "Percentage Rates",
			args: args{
				&rate,
				want: fees.TieredFees{
					[]fees.TierFee{
						*fees.NewTierFee(0.0, fees.NewPercentageFee(10)),
						*fees.NewTierFee(0.0, fees.NewPercentageFee(10)),
					},
				},
				wantErr: true,
			},
		},
		for _, tt := range tests{
		t.Run(tt.name, func (t *testing.T){
		got, err := FromProtoTieredRate(tt.args.proto)
		if (err != nil) != tt.wantErr{
		t.Errorf("FromProtoTieredRate() error = %v, wantErr %v", err, tt.wantErr)
		return
	}
		if !reflect.DeepEqual(got, tt.want){
		t.Errorf("FromProtoTieredRate() got = %v, want %v", got, tt.want)
	}
	})
	}
	}
