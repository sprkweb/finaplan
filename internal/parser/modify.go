package parser

import (
	"github.com/spf13/cobra"
	"github.com/sprkweb/finaplan-cli/pkg/finaplan"
)

type WrappedFunc func(plan *finaplan.FinancialPlan, args []string) error
type WrapperFunc func(cmd *cobra.Command, args []string)

func ModifyPlan(f WrappedFunc) WrapperFunc {
	return func(cmd *cobra.Command, args []string) {
		plan, err := ParsePlanFromStdin()
		if err != nil {
			panic(err)
		}

		err = f(plan, args)
		if err != nil {
			panic(err)
		}

		if err = PrintPlanToStdout(plan); err != nil {
			panic(err)
		}
	}
}
