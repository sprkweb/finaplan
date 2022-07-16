package finaplan

func Init(config *PlanConfig, intervalAmount uint64) *FinancialPlan {
	projection := make(Projection, intervalAmount)
	return &FinancialPlan{
		Config:     config,
		Projection: projection,
	}
}
