package finaplan

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestFinancialPlan_Invest(t *testing.T) {
	type args struct {
		interest  string
		intervals uint32
		start     uint32
		compound  bool
	}
	tests := []struct {
		name    string
		init    func() *FinancialPlan
		args    args
		want    []string
		wantErr assert.ErrorAssertionFunc
	}{
        {
        	name: "simple",
        	init: func() *FinancialPlan {
                plan := Init(DefaultConfig(), 7)
                plan.Add("300", 0, 0)
                return plan
        	},
        	args: args{
        		interest:  "10%",
        		intervals: 2,
        		start:     2,
        		compound:  false,
        	},
        	want: []string{"300", "300", "300", "315", "330", "345", "360"},
        	wantErr: assert.NoError,
        },
        {
        	name: "with add",
        	init: func() *FinancialPlan {
                plan := Init(DefaultConfig(), 6)
                plan.Add("300", 1, 0)
                return plan
        	},
        	args: args{
        		interest:  "0.1", // 10%
        		intervals: 2,
        		start:     2,
        		compound:  false,
        	},
        	want: []string{"300", "600", "900", "1245", "1605", "1980"},
        	wantErr: assert.NoError,
        },
        {
            name: "with one number",
        	init: func() *FinancialPlan {
                plan := Init(DefaultConfig(), 1)
                plan.Add("300", 1, 0)
                return plan
        	},
        	args: args{
        		interest:  "10%",
        		intervals: 2,
        		start:     2,
        		compound:  false,
        	},
        	want: []string{"300"},
        	wantErr: assert.NoError,
        },
        {
            name: "with no numbers",
        	init: func() *FinancialPlan {
                plan := Init(DefaultConfig(), 0)
                plan.Add("300", 1, 0)
                return plan
        	},
        	args: args{
        		interest:  "10%",
        		intervals: 2,
        		start:     2,
        		compound:  false,
        	},
        	want: []string{},
        	wantErr: assert.NoError,
        },
        {
            name: "compound",
        	init: func() *FinancialPlan {
                plan := Init(DefaultConfig(), 5)
                plan.Add("300", 0, 0)
                return plan
        	},
        	args: args{
        		interest:  "21%",
        		intervals: 2,
        		start:     0,
        		compound:  true,
        	},
        	want: []string{"300", "330", "363", "399.3", "439.23"},
        	wantErr: assert.NoError,
        },
        {
            name: "compound with add",
        	init: func() *FinancialPlan {
                plan := Init(DefaultConfig(), 9)
                plan.Add("200", 3, 2)
                return plan
        	},
        	args: args{
        		interest:  "0.1",
        		intervals: 1,
        		start:     1,
        		compound:  true,
        	},
        	want: []string{"0", "0", "200", "220", "242", "466.2", "512.82", "564.102", "820.5122"},
        	wantErr: assert.NoError,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plan := tt.init()
			tt.wantErr(t, plan.Invest(tt.args.interest, tt.args.intervals, tt.args.start, tt.args.compound))
			if len(plan.Projection) != len(tt.want) {
				t.Errorf("Got slice of unexpected length. Expected %d, got %d", len(tt.want), len(plan.Projection))
				return
			}
			for i, amount := range plan.Projection {
				if !amount.Equal(decimal.RequireFromString(tt.want[i])) {
					t.Errorf("Element №%d = %v does not match the expected value (%v)", i, amount, tt.want[i])
				}
			}
		})
	}
}
