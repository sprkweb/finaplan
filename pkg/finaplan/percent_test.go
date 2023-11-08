package finaplan

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func Test_percent(t *testing.T) {
	tests := []struct {
		give    string
		want    decimal.Decimal
		wantErr bool
	}{
		{
			give: "4%",
			want: decimal.RequireFromString("0.04"),
		},
		{
			give: "-215.51%",
			want: decimal.RequireFromString("-2.1551"),
		},
		{
			give: "0%",
			want: decimal.RequireFromString("0.00"),
		},
		{
			give: "0",
			want: decimal.RequireFromString("0"),
		},
		{
			give: "1",
			want: decimal.RequireFromString("1"),
		},
		{
			give: "4.9",
			want: decimal.RequireFromString("4.9"),
		},
		{
			give: "-0.01",
			want: decimal.RequireFromString("-0.01"),
		},
		{
			give:    "abc",
			want:    decimal.Decimal{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := percent("myparam", tt.give)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePercent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePercent() = %v, want %v", got, tt.want)
			}
		})
	}
}
