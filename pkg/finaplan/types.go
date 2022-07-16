package finaplan

import (
	"errors"
	"fmt"
)

type FinancialPlan struct {
	Config     *PlanConfig
	Projection Projection
}

type ProjectionUnit float64
type Projection []ProjectionUnit

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

func (c *PlanConfig) Validate() error {
	if c.IntervalLength < 1 {
		return errors.New("IntervalLength must be 1 or bigger")
	}
	switch c.IntervalType {
	case Days:
	case Weeks:
	case Months:
	case Years:
	default:
		return fmt.Errorf("incorrect IntervalType: %s", c.IntervalType)
	}
	return nil
}
