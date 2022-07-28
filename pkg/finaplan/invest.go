package finaplan

import (
	"errors"
	"math"
)

// Invest your capital with given `interest` per `intervals` starting after `start` intervals.
func (p *FinancialPlan) Invest(interest float64, intervals uint64, start uint64, compound bool) error {
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

func (p *FinancialPlan) calculateSimpleInterest(newProjection Projection, interest float64, intervals uint64, start uint64) {
	interestPerInterval := ProjectionUnit((interest - 1) / float64(intervals))

	var interestSum ProjectionUnit

	for i := start + 1; i <= uint64(len(p.Projection)-1); i++ {
		interestSum += p.Projection[i-1] * interestPerInterval
		newProjection[i] += interestSum
	}
}

func (p *FinancialPlan) calculateCompoundInterest(newProjection Projection, interest float64, intervals uint64, start uint64) {
	interestPerInterval := ProjectionUnit(math.Pow(interest, 1/float64(intervals)) - 1)

	for i := start + 1; i <= uint64(len(p.Projection)-1); i++ {
		newProjection[i] = newProjection[i-1]*(interestPerInterval+1) + (p.Projection[i] - p.Projection[i-1])
	}
}
