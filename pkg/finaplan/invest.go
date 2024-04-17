package finaplan

import (
	"github.com/shopspring/decimal"
)

// Invest your capital with given `interest` per `intervals` starting after `start` intervals.
func (p *FinancialPlan) Invest(interest string, intervals uint32, start uint32, compound bool) error {
	interestDecimal, err := percent("interest", interest)
	if err != nil {
		return err
	}
	if intervals < 1 {
		return ErrIntervalsLessThanOne
	}
	if int(start) > len(p.Projection)-1 {
		return nil
	}

	newProjection := make(Projection, len(p.Projection))
	copy(newProjection, p.Projection)

	if compound {
		p.calculateCompoundInterest(newProjection, interestDecimal, intervals, start)
	} else {
		p.calculateSimpleInterest(newProjection, interestDecimal, intervals, start)
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
	one := decimal.NewFromInt(1)
	interestBase := interest.Add(one)
	interestExpontent := one.Div(decimal.NewFromInt(int64(intervals)))
	interestPerInterval := interestBase.Pow(interestExpontent)

	for i := start + 1; i <= uint32(len(p.Projection)-1); i++ {
		newProjection[i] = newProjection[i-1].Mul(interestPerInterval).Add(p.Projection[i]).Sub(p.Projection[i-1])
	}
}
