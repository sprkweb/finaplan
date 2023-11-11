package finaplanio

import (
	"fmt"
	"io"
	"os"

	"github.com/sprkweb/finaplan/pkg/finaplan"
)

func PrintPlan(w io.Writer, p *finaplan.FinancialPlan) error {
	data, err := p.MarshalJSON()
	if err != nil {
		return fmt.Errorf("marshaling plan: %w", err)
	}

	if _, err = w.Write(data); err != nil {
		return fmt.Errorf("writing: %w", err)
	}
	return nil
}

func PrintPlanToStdout(plan *finaplan.FinancialPlan) error {
	return PrintPlan(os.Stdout, plan)
}
