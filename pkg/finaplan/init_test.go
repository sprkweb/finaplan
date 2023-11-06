package finaplan

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	tests := []uint32{5, 123, 0, 1}
	var plan *FinancialPlan
	for _, intervalAmount := range tests {
		t.Run(fmt.Sprintf("Init %d intervals", intervalAmount), func(t *testing.T) {
			plan = Init(DefaultConfig(), intervalAmount)
			if len(plan.Projection) != int(intervalAmount) {
				t.Errorf("Tried to initialize plan with %d intervals, but got %d", intervalAmount, len(plan.Projection))
			}
		})
	}
}
