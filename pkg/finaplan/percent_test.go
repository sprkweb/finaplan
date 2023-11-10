package finaplan

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_percent(t *testing.T) {
	tests := []struct {
		give    string
		want    decimal.Decimal
		wantErr assert.ErrorAssertionFunc
	}{
		{
			give:    "4%",
			want:    decimal.RequireFromString("0.04"),
			wantErr: assert.NoError,
		},
		{
			give:    "-215.51%",
			want:    decimal.RequireFromString("-2.1551"),
			wantErr: assert.NoError,
		},
		{
			give:    "0%",
			want:    decimal.RequireFromString("0.00"),
			wantErr: assert.NoError,
		},
		{
			give:    "0",
			want:    decimal.RequireFromString("0"),
			wantErr: assert.NoError,
		},
		{
			give:    "1",
			want:    decimal.RequireFromString("1"),
			wantErr: assert.NoError,
		},
		{
			give:    "4.9",
			want:    decimal.RequireFromString("4.9"),
			wantErr: assert.NoError,
		},
		{
			give:    "-0.01",
			want:    decimal.RequireFromString("-0.01"),
			wantErr: assert.NoError,
		},
		{
			give:    "abc",
			want:    decimal.Decimal{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got, err := percent("myparam", tt.give)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
