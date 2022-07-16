package finaplan

func (p FinancialPlan) Add(amount ProjectionUnit, each uint64, start uint64) {
	maxIndex := uint64(len(p.Projection)) - 1
	if start > maxIndex {
		return
	}

	for i := start; i <= maxIndex; i++ {
		if each >= 1 {
			intervalsPassed := (i - start) / each
			p.Projection[i] += amount * ProjectionUnit(intervalsPassed+1)
		} else {
			p.Projection[i] += amount
		}
	}
}
