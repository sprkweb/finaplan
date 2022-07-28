package finaplan

import (
	"errors"
	"fmt"
)

// FinancialPlan is the main structure with information about the financial plan.
// It is the output of any FinaPlan module.
type FinancialPlan struct {
	// Configuration of your plan
	Config *PlanConfig
	// The plan itself; the projected capital changing over the time
	Projection Projection
}

type ProjectionUnit float64
type Projection []ProjectionUnit

// PlanConfig is configuration of your plan
type PlanConfig struct {
	// The indivisible unit of time measurement in the plan.
	//
	// "days" | "weeks" | "months" | "years"
	IntervalType IntervalType `yaml:"interval_type"`
	// Amount of the time units in 1 interval
	//
	// For example,
	// if you want to calculate your capital for each 3 months, it is 3.
	IntervalLength uint32 `yaml:"interval_length"`
}

type IntervalType string

// IntervalType is a type of indivisible units of time measurement in the plan.
const (
	Days   IntervalType = "days"
	Weeks  IntervalType = "weeks"
	Months IntervalType = "months"
	Years  IntervalType = "years"
)

// GetIntervalUnit gets the name of the given unit (singular form, lower case).
//
// Error is returned if the given IntervalType does not exist.
func GetIntervalUnit(t IntervalType) (string, error) {
	switch t {
	case Days:
		return "day", nil
	case Weeks:
		return "week", nil
	case Months:
		return "month", nil
	case Years:
		return "year", nil
	default:
		return "", fmt.Errorf("wrong interval type: %s", t)
	}
}

func DefaultConfig() *PlanConfig {
	return &PlanConfig{
		IntervalType:   Months,
		IntervalLength: 1,
	}
}

// Validate PlanConfig. If it is valid, nil is returned
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
