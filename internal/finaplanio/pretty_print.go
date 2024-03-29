package finaplanio

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sprkweb/finaplan/pkg/finaplan"
)

func PrettyPrint(w io.Writer, plan *finaplan.FinancialPlan) error {
	projection := plan.PrettyPrint()
	if len(projection) < 1 {
		return nil
	}

	unit := plan.Config.IntervalType.GetIntervalUnit()
	unit = strings.ToUpper(unit[:1]) + unit[1:] // capitalize

	for i, v := range projection {
		interval := uint64(plan.Config.IntervalLength) * uint64(i)
		fmt.Fprintf(w, "%s %d  \t| %s\n", unit, interval, v)
	}

	return nil
}

func PrettyPrintPlanToStdout(plan *finaplan.FinancialPlan) error {
	return PrettyPrint(os.Stdout, plan)
}
