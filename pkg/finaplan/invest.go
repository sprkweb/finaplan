package finaplan

import (
	"errors"
	"math"

	"github.com/shopspring/decimal"
)

// Invest your capital with given `interest` per `intervals` starting after `start` intervals.
func (p *FinancialPlan) Invest(interest decimal.Decimal, intervals uint32, start uint32, compound bool) error {
	if int(start) > len(p.Projection)-1 {
		return nil
	}
	if intervals < 1 {
		return errors.New("intervals must be greater than 1")
	}

	newProjection := make(Projection, len(p.Projection))
	copy(newProjection, p.Projection)

	if compound {
		p.calculateCompoundInterest(newProjection, interest, intervals, start)
	} else {
		p.calculateSimpleInterest(newProjection, interest, intervals, start)
	}
	p.Projection = newProjection
	return nil
}

func (p *FinancialPlan) calculateSimpleInterest(newProjection Projection, interestPercent decimal.Decimal, intervals uint32, start uint32) {
	// interestPerInterval = interestPercent / intervals
	interestPerInterval := interestPercent.Div(decimal.NewFromInt(int64(intervals)))

	var interestSum decimal.Decimal

	for i := start + 1; i <= uint32(len(p.Projection)-1); i++ {
		// interestSum += projection[i-1] * interestPerInterval
		interestSum = interestSum.Add(p.Projection[i-1].Mul(interestPerInterval))
		// newProjection[i] += interestSum
		newProjection[i] = newProjection[i].Add(interestSum)
	}
}

func (p *FinancialPlan) calculateCompoundInterest(newProjection Projection, interest decimal.Decimal, intervals uint32, start uint32) {
	// interestPerInterval := (interest + 1) ^ (1 / intervals)
	// decimal package does not support neither root nor power to non-integer numbers
	// so we have to convert to float64 and use standard pow function here
	one := decimal.NewFromInt(1)
	interestBase := interest.Add(one).InexactFloat64()
	interestExpontent := one.Div(decimal.NewFromInt(int64(intervals))).InexactFloat64()
	interestPerInterval := decimal.NewFromFloat(math.Pow(interestBase, interestExpontent))

	for i := start + 1; i <= uint32(len(p.Projection)-1); i++ {
		newProjection[i] = newProjection[i-1].Mul(interestPerInterval).Add(p.Projection[i]).Sub(p.Projection[i-1])
	}
}
