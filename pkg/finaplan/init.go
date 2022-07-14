package finaplan

func Init(config *PlanConfig, intervalAmount uint64) *FinancialPlan {
	projection := make(projection, intervalAmount)
	return &FinancialPlan{
		Config:     config,
		Projection: projection,
	}
}
