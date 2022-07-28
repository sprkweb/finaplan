package finaplan

// Add a certain amount of money to your financial plan (e.g. savings)
//
// regularly (`each` N intervals) or once (`each` 0 intervals)
//
// starting after `start` intervals
func (p FinancialPlan) Add(amount ProjectionUnit, each uint64, start uint64) {
	var maxIndex int = len(p.Projection) - 1
	if int(start) > maxIndex {
		return
	}

	for i := start; i <= uint64(maxIndex); i++ {
		if each >= 1 {
			intervalsPassed := (i - start) / each
			p.Projection[i] += amount * ProjectionUnit(intervalsPassed+1)
		} else {
			p.Projection[i] += amount
		}
	}
}
