package mapping

import (
	"FeeCalculatorService/fees"
	proto "FeeCalculatorService/gen/guido4f.fee"
	"reflect"
	"testing"
)

func TestFromProtoTieredRate(t *testing.T) {
	type args struct {
		proto *proto.FeeCalculatorRequest_TieredRate
	}

	request_tieredRate := buildTieredRateProto(
		buildTier(20, 10),
		buildTier(30, 15))

	tests := []struct {
		name    string
		args    args
		want    fees.TieredFees
		wantErr bool
	}{
		{name: "Percentage Rates",
			args: args{
				&request_tieredRate,
			},
			want: fees.TieredFees{
				[]fees.TierFee{
					*fees.NewTierFee(20.0, fees.NewPercentageFee(10)),
					*fees.NewTierFee(30.0, fees.NewPercentageFee(15)),
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromProtoTieredRate(tt.args.proto)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromProtoTieredRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromProtoTieredRate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildTier(bound float64, percentage float64) *proto.Tier {
	return &proto.Tier{
		HighBound: &bound,
		Rate: &proto.Rate{
			Rate: &proto.Rate_Percentage{percentage},
		},
	}
}

func buildTieredRateProto(tiers ...*proto.Tier) proto.FeeCalculatorRequest_TieredRate {

	tiersRates := proto.TieredRates{Tiers: tiers}
	request_tieredRate := proto.FeeCalculatorRequest_TieredRate{&tiersRates}
	return request_tieredRate
}
