package finaplan

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

// FinancialPlan is the main structure with information about the financial plan.
// It is the output of any FinaPlan module.
type FinancialPlan struct {
	// Configuration of your plan
	Config *PlanConfig
	// The plan itself; the projected capital changing over the time
	Projection Projection
}

type Projection []decimal.Decimal

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

// String returns the name of the given unit (multiple form, lower case)
func (t IntervalType) String() string {
	return string(t)
}

// GetIntervalUnit gets the name of the given unit (singular form, lower case).
func (t IntervalType) GetIntervalUnit() string {
	switch t {
	case Days:
		return "day"
	case Weeks:
		return "week"
	case Months:
		return "month"
	case Years:
		return "year"
	}
	return ""
}

func (t IntervalType) Validate() error {
	switch t {
	case Days,
		Weeks,
		Months,
		Years:
		return nil
	default:
		return fmt.Errorf("wrong interval type: %q", t)
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
	var errs []error
	if c.IntervalLength < 1 {
		errs = append(errs, errors.New("IntervalLength must be 1 or bigger"))
	}
	if err := c.IntervalType.Validate(); err != nil {
		errs = append(errs, fmt.Errorf("invalid IntervalType: %w", err))
	}
	return errors.Join(errs...)
}
