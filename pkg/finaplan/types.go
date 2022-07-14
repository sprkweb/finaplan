package finaplan

type FinancialPlan struct {
	Config     *PlanConfig
	Projection projection
}

type projection []int64

type PlanConfig struct {
	IntervalType   IntervalType
	IntervalLength uint32
}

type IntervalType uint8

const (
	Days IntervalType = iota
	Weeks
	Months
	Years
)

func DefaultConfig() *PlanConfig {
	return &PlanConfig{
		IntervalType:   Days,
		IntervalLength: 1,
	}
}
