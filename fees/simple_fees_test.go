package fees

import "testing"

func TestCalculate(t *testing.T) {
	type args struct {
		rate        FeeRate
		marketValue float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"100%", args{PercentageFee{100.0}, 120000}, 120000},
		{"50%", args{PercentageFee{50.0}, 120000}, 60000},
		{"0%", args{PercentageFee{0.0}, 120000}, 0},
		{"100bps", args{BasisPointFee{100.0}, 120000}, 1200},
		{"50bps", args{BasisPointFee{50.0}, 120000}, 600},
		{"0bps", args{BasisPointFee{0.0}, 120000}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.rate, tt.args.marketValue); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
