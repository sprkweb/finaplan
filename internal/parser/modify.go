package parser

import (
	"github.com/spf13/cobra"
	"github.com/sprkweb/finaplan/pkg/finaplan"
)

type WrappedFunc func(plan *finaplan.FinancialPlan, args []string) error
type WrapperFunc func(cmd *cobra.Command, args []string)

// ModifyPlan wraps your command handler function with input / output handlers:
//
// it parses financial plan from stdin and then, after your function modifies it, the plan is printed to stdout.
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
