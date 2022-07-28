package finaplan

// Init an empty financial plan with given `config` for `intervalAmount` intervals
func Init(config *PlanConfig, intervalAmount uint64) *FinancialPlan {
	projection := make(Projection, intervalAmount)
	return &FinancialPlan{
		Config:     config,
		Projection: projection,
	}
}
