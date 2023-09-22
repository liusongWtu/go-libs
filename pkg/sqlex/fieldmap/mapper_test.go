package fieldmap

import (
	"reflect"
	"testing"
)

func TestToLowerCamel(t *testing.T) {
	type args struct {
		fields []string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "1",
			args: args{
				fields: []string{"create_time", "payee", "cost_time"},
			},
			want: map[string]string{
				"createTime": "create_time",
				"payee":      "payee",
				"costTime":   "cost_time",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLowerCamel(tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
