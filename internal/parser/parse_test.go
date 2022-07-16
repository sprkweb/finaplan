package parser

import (
	"fmt"
	"github.com/sprkweb/finaplan-cli/pkg/finaplan"
	"testing"
)

func TestParsePlanWithGeneratedInput(t *testing.T) {
	initialPlan := finaplan.Init(finaplan.DefaultConfig(), 12)
	output, err := PrintPlan(initialPlan)
	if err != nil {
		t.Errorf("Error printing plan: %s", err)
	}
	parsedPlan, err := ParsePlan(output)
	if err != nil {
		t.Errorf("Error parsing plan: %s", err)
	}
	if len(parsedPlan.Projection) != len(initialPlan.Projection) {
		t.Errorf("Mismatched length of the parsed plan: %d instead of %d", len(parsedPlan.Projection), len(initialPlan.Projection))
	}
	if parsedPlan.Config.IntervalType != initialPlan.Config.IntervalType ||
		parsedPlan.Config.IntervalLength != initialPlan.Config.IntervalLength {
		t.Errorf("Mismatched parsed config")
	}
}
func TestParsePlanWithCorrectInput(t *testing.T) {
	rightInputs := []string{
		"---\ninterval_type: weeks\ninterval_length: 1\n---\n",
		"---\ninterval_type: days\ninterval_length: 123\n---\n1\n0\n0",
		"---\ninterval_type: months\ninterval_length: 2\n---\n1\n0\n0\n",
		"---\ninterval_length: 9\ninterval_type: months\n---\n2\n",
	}
	for i, input := range rightInputs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			_, err := ParsePlan(input)
			if err != nil {
				t.Errorf("Got error with correct input: %s", err)
			}
		})
	}
}

func TestParsePlanWithIncorrectInput(t *testing.T) {
	incorrectInputs := []string{
		"1\n2\n3",
		"---\nfoo\n---\n1",
		"---\ninterval_length: 9\ninterval_type: bars\n---\n2\n",
		"---\ninterval_type: weeks\ninterval_length: -1\n---\n1",
		"---\ninterval_type: years\ninterval_length: 0\n---\n1\n0\n0",
		"interval_type: years\ninterval_length: 2\n---\n1\n0\n0",
		"---\ninterval_type: months\ninterval_length: 3\n1\n0\n0",
	}
	for i, input := range incorrectInputs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result, err := ParsePlan(input)
			if err == nil {
				t.Errorf("Got no error with incorrect input. Parsed result: %v", *result)
			}
		})
	}
}
