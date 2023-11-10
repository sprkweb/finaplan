package finaplan

// Print returns resulting projection of your capital changing over time according to the plan
//
// The output is raw numbers which are intended for usage by programs or saving
func (p *FinancialPlan) Print() []string {
	output := make([]string, 0, len(p.Projection))
	for _, num := range p.Projection {
		output = append(output, num.String())
	}
	return output
}

// PrettyPrint returns resulting projection of your capital changing over time according to the plan
//
// The output is formatted for humans, so it might be not as precise as Print
func (p *FinancialPlan) PrettyPrint() []string {
	output := make([]string, 0, len(p.Projection))
	for _, num := range p.Projection {
		output = append(output, num.StringFixedBank(2))
	}
	return output
}
