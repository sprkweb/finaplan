package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sprkweb/finaplan/internal/parser"
	"github.com/sprkweb/finaplan/pkg/finaplan"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add AMOUNT",
	Short: "Add a certain amount of money to your plan",
	Long: `Add a certain amount of money to your financial plan (e.g. savings) once or regularly

Example:
$ finaplan init | finaplan add 300 --each 2
---
interval_type: days
interval_length: 1
---
300
300
600
600
900`,
	Args: cobra.ExactArgs(1),
	Run: parser.ModifyPlan(func(plan *finaplan.FinancialPlan, args []string) error {
		return plan.Add(args[0], AddEach, AddStart)
	}),
}

var AddEach uint32
var AddStart uint32

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().Uint32Var(&AddEach, "each", 0, "period; number of intervals after which the same amount is added (0 if you do not want it to repeat)")
	addCmd.Flags().Uint32Var(&AddStart, "start", 0, "when do you want to add the amount (0 = at the very beginning of the plan)")
}
