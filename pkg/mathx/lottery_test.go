package mathx

import "testing"

func TestLottery(t *testing.T) {
	weights := []int{0, 1, 0, 2}

	for i := 0; i < 100; i++ {
		index := Lottery(weights)
		t.Log(index)
	}
}

func TestLotteryByRate(t *testing.T) {
	type args struct {
		rate      float64
		precision float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				rate:      1,
				precision: 4,
			},
			want: true,
		},
		{
			name: "0.5",
			args: args{
				rate:      0.5,
				precision: 4,
			},
			want: true,
		},
		{
			name: "0",
			args: args{
				rate:      0,
				precision: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LotteryByRate(tt.args.rate, tt.args.precision); got != tt.want {
				t.Errorf("LotteryByRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
