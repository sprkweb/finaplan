package parser

import (
	"github.com/sprkweb/finaplan-cli/pkg/finaplan"
	"testing"
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
	var intervalAmount uint64 = 4
	plan := finaplan.Init(config, intervalAmount)

	result, err := PrintPlan(plan)
	if err != nil {
		t.Errorf("Got error: %s", err)
	}
	if result != expected {
		t.Errorf("Got wrong output:\n%s", result)
	}
}
