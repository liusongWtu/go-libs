package stringx

import "testing"

func TestFloatFormatByDecimalPlace(t *testing.T) {
	type args struct {
		value         float64
		decimalPlaces int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "无小数位-保留2位小数",
			args: args{
				value:         1,
				decimalPlaces: 2,
			},
			want: "1",
		},
		{
			name: "无小数位-保留0位小数",
			args: args{
				value:         1,
				decimalPlaces: 0,
			},
			want: "1",
		},
		{
			name: "有小数位-保留0位小数",
			args: args{
				value:         1.235,
				decimalPlaces: 0,
			},
			want: "1",
		},
		{
			name: "有小数位-保留2位小数",
			args: args{
				value:         1.235,
				decimalPlaces: 2,
			},
			want: "1.24",
		},
		{
			name: "有小数位-保留3位小数",
			args: args{
				value:         1.235,
				decimalPlaces: 3,
			},
			want: "1.235",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatFormatByDecimalPlace(tt.args.value, tt.args.decimalPlaces); got != tt.want {
				t.Errorf("FloatFormatByDecimalPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}
