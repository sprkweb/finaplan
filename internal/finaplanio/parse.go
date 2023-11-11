package finaplanio

import (
	"fmt"
	"io"
	"os"

	"github.com/sprkweb/finaplan/pkg/finaplan"
)

func ParsePlan(r io.Reader) (*finaplan.FinancialPlan, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("reading: %w", err)
	}

	var plan finaplan.FinancialPlan
	if err := plan.UnmarshalJSON(data); err != nil {
		return nil, fmt.Errorf("unmarshaling plan: %w", err)
	}
	return &plan, nil
}

func ParsePlanFromStdin() (*finaplan.FinancialPlan, error) {
	return ParsePlan(os.Stdin)
}
