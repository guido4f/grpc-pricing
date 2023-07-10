package fees

import (
	"reflect"
	"testing"
)

func TestTieredFees_calculateTierFee(t *testing.T) {
	type fields struct {
		Tiers []TierFee
	}
	type args struct {
		amount float64
	}
	fee1 := TierFee{
		1300,
		percentageFee{10},
	}
	fee2 := TierFee{
		500,
		percentageFee{5},
	}
	fee3 := TierFee{
		100,
		percentageFee{20},
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   CalculatedTieredFees
	}{
		{
			"Single Tiers",
			fields{[]TierFee{fee1}},
			args{2000.0},
			CalculatedTieredFees{[]CalculatedTierFee{*buildPctCalculaterTier(1300.0, 10.0, 1300.0, 130.0)},
				130.0},
		},
		{
			"Test Based Descending supplied Tiers",
			fields{[]TierFee{fee1, fee2, fee3}},
			args{2000.0}, CalculatedTieredFees{[]CalculatedTierFee{
				*buildPctCalculaterTier(100.0, 20.0, 100.0, 20.0),
				*buildPctCalculaterTier(500.0, 5, 400.0, 20.0),
				*buildPctCalculaterTier(1300.0, 10.0, 800.0, 80.0),
			},
				80 + 20 + 20},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := TieredFees{
				Tiers: tt.fields.Tiers,
			}
			if got := r.calculateTierFee(tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateTierFee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildPctCalculaterTier(bound float64, fee float64, tier float64, feeAmt float64) *CalculatedTierFee {
	return NewCalculatedTierFee(
		*NewTierFee(bound, *NewPercentageFee(fee)),
		tier,
		feeAmt,
	)
}
