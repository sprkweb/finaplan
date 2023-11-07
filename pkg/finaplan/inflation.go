package finaplan

import (
	"errors"
	"math"

	"github.com/shopspring/decimal"
)

// Adjust your capital growth over time for inflation.
//
// All the sums will be adjusted to the prices of `start = 0`,
// considering that the prices grow by `inflation` percent per `intervals`
//
// This should be applied after all the other calculations,
// right before printing the final result.
func (p *FinancialPlan) Inflation(inflation decimal.Decimal, intervals uint32) error {
	if intervals < 1 {
		return errors.New("intervals must be equal or greater than 1")
	}

	// inflationPerInterval := (inflation + 1) ^ (1 / intervals)
	// decimal package does not support neither root nor power to non-integer numbers
	// so we have to convert to float64 and use standard pow function here
	one := decimal.NewFromInt(1)
	base := inflation.Add(one).InexactFloat64()
	expontent := one.Div(decimal.NewFromInt(int64(intervals))).InexactFloat64()
	inflationPerInterval := decimal.NewFromFloat(math.Pow(base, expontent))

	for i := 0; i < len(p.Projection); i++ {
		inflationSoFar := inflationPerInterval.Pow(decimal.NewFromInt(int64(i)))
		p.Projection[i] = p.Projection[i].Div(inflationSoFar)
	}
	return nil
}
