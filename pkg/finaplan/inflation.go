package finaplan

import (
	"github.com/shopspring/decimal"
)

// Adjust your capital growth over time for inflation.
//
// All the sums will be adjusted to the prices of `start = 0`,
// considering that the prices grow by `inflation` percent per `intervals`
//
// This should be applied after all the other calculations,
// right before printing the final result.
func (p *FinancialPlan) Inflation(inflation string, intervals uint32) error {
	inflationDecimal, err := percent("inflation", inflation)
	if err != nil {
		return err
	}

	if intervals < 1 {
		return ErrIntervalsLessThanOne
	}

	// inflationPerInterval := (inflation + 1) ^ (1 / intervals)
	one := decimal.NewFromInt(1)
	base := inflationDecimal.Add(one)
	expontent := one.Div(decimal.NewFromInt(int64(intervals)))
	inflationPerInterval := base.Pow(expontent)

	for i := 0; i < len(p.Projection); i++ {
		inflationSoFar := inflationPerInterval.Pow(decimal.NewFromInt(int64(i)))
		p.Projection[i] = p.Projection[i].Div(inflationSoFar)
	}
	return nil
}
