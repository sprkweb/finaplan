package finaplan

type FinancialPlan struct {
	Config     *PlanConfig
	Projection projection
}

type projection []float64

type PlanConfig struct {
	IntervalType   IntervalType `yaml:"interval_type"`
	IntervalLength uint32       `yaml:"interval_length"`
}

type IntervalType string

const (
	Days   IntervalType = "days"
	Weeks  IntervalType = "weeks"
	Months IntervalType = "months"
	Years  IntervalType = "years"
)

func DefaultConfig() *PlanConfig {
	return &PlanConfig{
		IntervalType:   Days,
		IntervalLength: 1,
	}
}
