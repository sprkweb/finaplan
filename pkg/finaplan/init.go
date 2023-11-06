package finaplan

import "github.com/shopspring/decimal"

// Init an empty financial plan with given `config` for `intervalAmount` intervals
func Init(config *PlanConfig, intervalAmount uint32) *FinancialPlan {
	decimal.DivisionPrecision = 32
	projection := make(Projection, intervalAmount)
	return &FinancialPlan{
		Config:     config,
		Projection: projection,
	}
}
