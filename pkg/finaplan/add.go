package finaplan

import (
	"github.com/shopspring/decimal"
)

// Add a certain amount of money to your financial plan (e.g. savings)
//
// regularly (`each` N intervals) or once (`each` 0 intervals)
//
// starting after `start` intervals
func (p FinancialPlan) Add(amount decimal.Decimal, each uint32, start uint32) {
	var maxIndex int = len(p.Projection) - 1
	if int(start) > maxIndex {
		return
	}

	if each >= 1 {
		for i := start; i <= uint32(maxIndex); i++ {
			// intervalsPassed = 1 + (i - start) / each
			intervalsPassed := decimal.NewFromInt(1 + int64((i-start)/each))
			// projection += amount * intervalsPassed
			p.Projection[i] = p.Projection[i].Add(amount.Mul(intervalsPassed))
		}
	} else {
		for i := start; i <= uint32(maxIndex); i++ {
			// projection += amount
			p.Projection[i] = p.Projection[i].Add(amount)
		}
	}
}
