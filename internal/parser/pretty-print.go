package parser

import (
	"fmt"
	"github.com/sprkweb/finaplan-cli/pkg/finaplan"
	"strings"
)

func PrettyPrintPlanToStdout(plan *finaplan.FinancialPlan) error {
	unit, err := finaplan.GetIntervalUnit(plan.Config.IntervalType)
	if err != nil {
		return err
	}

	// capitalize
	unit = fmt.Sprintf("%s%s", strings.ToUpper(unit[0:1]), unit[1:])

	labelLength := len(sprintLabel(unit, uint64(len(plan.Projection)-1)))
	format := fmt.Sprintf("%%%ds | %%v\n", labelLength)
	for i, v := range plan.Projection {
		interval := uint64(plan.Config.IntervalLength) * uint64(i)
		label := sprintLabel(unit, interval)
		fmt.Printf(format, label, v)
	}

	return nil
}

func sprintLabel(unit string, interval uint64) string {
	return fmt.Sprintf("%s %d", unit, interval)
}
