package finaplan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// right now some of these test will not pass
// because of float64 calculation inaccuracies when taking root
//
// need to find a way to make these calculations in decimal numbers
func TestFinancialPlan_Inflation(t *testing.T) {
	tests := []struct {
		name      string
		init      func() *FinancialPlan
		inflation string
		intervals uint32
		want      []string
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "no capital",
			init: func() *FinancialPlan {
				return Init(DefaultConfig(), 7)
			},
			inflation: "0.21", // 21%
			intervals: 2,
			want:      []string{"0", "0", "0", "0", "0", "0", "0"},
			wantErr:   assert.NoError,
		},
		{
			name: "static sum with 300%",
			init: func() *FinancialPlan {
				plan := Init(DefaultConfig(), 7)
				_ = plan.Add("640", 0, 0)
				return plan
			},
			inflation: "3", // 300%
			intervals: 2,
			want:      []string{"640", "320", "160", "80", "40", "20", "10"},
			wantErr:   assert.NoError,
		},
		{
			name: "static sum with 33.1%",
			init: func() *FinancialPlan {
				plan := Init(DefaultConfig(), 7)
				_ = plan.Add("2679.61179333", 0, 0)
				return plan
			},
			inflation: "0.331", // 33.1%
			intervals: 3,
			want:      []string{"2679.61179333", "2436.01072121", "2214.5552011", "2013.232001", "1830.21091", "1663.8281", "1512.571"},
			wantErr:   assert.NoError,
		},
		{
			name: "periodic replenishment",
			init: func() *FinancialPlan {
				plan := Init(DefaultConfig(), 7)
				_ = plan.Add("640", 2, 2)
				return plan
			},
			inflation: "1", // 100%
			intervals: 1,
			want:      []string{"0", "0", "160", "80", "80", "40", "30"},
			wantErr:   assert.NoError,
		},
		{
			name: "same percent with investements",
			init: func() *FinancialPlan {
				plan := Init(DefaultConfig(), 7)
				_ = plan.Add("300", 0, 0)
				_ = plan.Invest("0.5", 1, 1, true)
				return plan
			},
			inflation: "0.5", // 21%
			intervals: 1,
			want:      []string{"300", "200", "200", "200", "200", "200", "200"},
			wantErr:   assert.NoError,
		},
		{
			name: "lower percent with investements",
			init: func() *FinancialPlan {
				plan := Init(DefaultConfig(), 7)
				_ = plan.Add("300", 0, 0)
				_ = plan.Invest("0.21", 1, 0, true)
				return plan
			},
			inflation: "0.1", // 10%
			intervals: 1,
			want:      []string{"300", "330", "363", "399.3", "439.23", "483.153", "531.4683"},
			wantErr:   assert.NoError,
		},
		{
			name: "deflation",
			init: func() *FinancialPlan {
				plan := Init(DefaultConfig(), 7)
				_ = plan.Add("300", 0, 0)
				return plan
			},
			inflation: "-0.01", // -1%
			intervals: 1,
			want:      []string{"300", "303", "306.03", "309.0903", "312.181203", "315.30301503", "318.4560451803"},
			wantErr:   assert.NoError,
		},
		{
			name: "zero inflation",
			init: func() *FinancialPlan {
				plan := Init(DefaultConfig(), 7)
				_ = plan.Add("300", 0, 0)
				return plan
			},
			inflation: "0", // 0%
			intervals: 1,
			want:      []string{"300", "300", "300", "300", "300", "300", "300"},
			wantErr:   assert.NoError,
		},
		{
			name: "per 0 intervals",
			init: func() *FinancialPlan {
				plan := Init(DefaultConfig(), 7)
				_ = plan.Add("300", 0, 0)
				return plan
			},
			inflation: "-0.01", // -1%
			intervals: 0,
			want:      []string{"300", "300", "300", "300", "300", "300", "300"},
			wantErr:   assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plan := tt.init()
			tt.wantErr(t, plan.Inflation(tt.inflation, tt.intervals))
			assert.Equal(t, tt.want, plan.Print())
		})
	}
}
