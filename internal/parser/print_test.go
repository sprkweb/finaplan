package parser

import (
	"testing"

	"github.com/sprkweb/finaplan-cli/finaplan/pkg/finaplan"
)

func TestPrintPlan(t *testing.T) {
	expected := `---
interval_type: months
interval_length: 2
---
0
0
0
0
`
	config := &finaplan.PlanConfig{
		IntervalType:   finaplan.Months,
		IntervalLength: 2,
	}
	intervalAmount := uint32(4)
	plan := finaplan.Init(config, intervalAmount)

	result, err := PrintPlan(plan)
	if err != nil {
		t.Errorf("Got error: %s", err)
	}
	if result != expected {
		t.Errorf("Got wrong output:\n%s", result)
	}
}
