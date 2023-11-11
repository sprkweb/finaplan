package finaplan

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestFinancialPlan_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		give    string
		want    *FinancialPlan
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "0 weeks",
			give: `{ "version": "0.1.0", "config": { "interval_type": "weeks", "interval_length": 1 }, "projection": [] }`,
			want: &FinancialPlan{
				Config: &PlanConfig{
					IntervalType:   Weeks,
					IntervalLength: 1,
				},
				Projection: []decimal.Decimal{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "5 days",
			give: `{ "version": "0.1.0", "config": { "interval_type": "days", "interval_length": 123 }, "projection": ["-1.512","0","0.95","420.1337","9"] }`,
			want: &FinancialPlan{
				Config: &PlanConfig{
					IntervalType:   Days,
					IntervalLength: 123,
				},
				Projection: []decimal.Decimal{
					decimal.RequireFromString("-1.512"),
					decimal.RequireFromString("0"),
					decimal.RequireFromString("0.95"),
					decimal.RequireFromString("420.1337"),
					decimal.RequireFromString("9"),
				},
			},
			wantErr: assert.NoError,
		},
		{
			name:    "wrong format version",
			give:    `{ "version": "4.2.0", "config": { "interval_type": "days", "interval_length": 123 }, "projection": ["-1.512","0","0.95","420.1337","9"] }`,
			want:    &FinancialPlan{},
			wantErr: assert.Error,
		},
		{
			name:    "zero interval length",
			give:    `{ "version": "0.1.0", "config": { "interval_type": "days", "interval_length": 0 }, "projection": ["-1.512","0","0.95","420.1337","9"] }`,
			want:    &FinancialPlan{},
			wantErr: assert.Error,
		},
		{
			name:    "negative interval length",
			give:    `{ "version": "0.1.0", "config": { "interval_type": "days", "interval_length": -1 }, "projection": ["-1.512","0","0.95","420.1337","9"] }`,
			want:    &FinancialPlan{},
			wantErr: assert.Error,
		},
		{
			name:    "invalid interval type",
			give:    `{ "version": "0.1.0", "config": { "interval_type": "decades", "interval_length": 123 }, "projection": ["-1.512","0","0.95","420.1337","9"] }`,
			want:    &FinancialPlan{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p FinancialPlan
			err := p.UnmarshalJSON([]byte(tt.give))
			assert.Equal(t, tt.want, &p)
			tt.wantErr(t, err)
		})
	}
}

func TestFinancialPlan_MarshalJSON(t *testing.T) {
	type fields struct {
		Config     *PlanConfig
		Projection Projection
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "four zeroes",
			fields: fields{
				Config: &PlanConfig{
					IntervalType:   Months,
					IntervalLength: 2,
				},
				Projection: []decimal.Decimal{
					decimal.NewFromInt(0),
					decimal.NewFromInt(0),
					decimal.NewFromInt(0),
					decimal.NewFromInt(0),
				},
			},
			want:    `{"version":"0.1.0","config":{"interval_type":"months","interval_length":2},"projection":["0","0","0","0"]}`,
			wantErr: assert.NoError,
		},
		{
			name: "changing sum",
			fields: fields{
				Config: &PlanConfig{
					IntervalType:   Days,
					IntervalLength: 1,
				},
				Projection: []decimal.Decimal{
					decimal.RequireFromString("-1.512"),
					decimal.RequireFromString("0"),
					decimal.RequireFromString("0.95"),
					decimal.RequireFromString("420.1337"),
					decimal.RequireFromString("9"),
				},
			},
			want:    `{"version":"0.1.0","config":{"interval_type":"days","interval_length":1},"projection":["-1.512","0","0.95","420.1337","9"]}`,
			wantErr: assert.NoError,
		},
		{
			name: "empty projection",
			fields: fields{
				Config: &PlanConfig{
					IntervalType:   Years,
					IntervalLength: 10,
				},
				Projection: []decimal.Decimal{},
			},
			want:    `{"version":"0.1.0","config":{"interval_type":"years","interval_length":10},"projection":[]}`,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &FinancialPlan{
				Config:     tt.fields.Config,
				Projection: tt.fields.Projection,
			}
			got, err := p.MarshalJSON()
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestFinancialPlan_Marshal_and_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		plan *FinancialPlan
	}{
		{
			name: "four zeroes",
			plan: &FinancialPlan{
				Config: &PlanConfig{
					IntervalType:   Months,
					IntervalLength: 2,
				},
				Projection: []decimal.Decimal{
					decimal.NewFromInt(0),
					decimal.NewFromInt(0),
					decimal.NewFromInt(0),
					decimal.NewFromInt(0),
				},
			},
		},
		{
			name: "changing sum",
			plan: &FinancialPlan{
				Config: &PlanConfig{
					IntervalType:   Days,
					IntervalLength: 1,
				},
				Projection: []decimal.Decimal{
					decimal.RequireFromString("-1.512"),
					decimal.RequireFromString("0"),
					decimal.RequireFromString("0.95"),
					decimal.RequireFromString("420.1337"),
					decimal.RequireFromString("9"),
				},
			},
		},
		{
			name: "empty projection",
			plan: &FinancialPlan{
				Config: &PlanConfig{
					IntervalType:   Years,
					IntervalLength: 10,
				},
				Projection: []decimal.Decimal{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := tt.plan.MarshalJSON()
			assert.NoError(t, err)

			var got FinancialPlan
			assert.NoError(t, got.UnmarshalJSON(encoded))
			assert.Equal(t, tt.plan, &got)
		})
	}
}
