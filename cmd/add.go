package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sprkweb/finaplan-cli/internal/parser"
	"github.com/sprkweb/finaplan-cli/pkg/finaplan"
	"strconv"
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
		amount, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return err
		}

		plan.Add(finaplan.ProjectionUnit(amount), AddEach, AddStart)
		return nil
	}),
}

var AddEach uint64
var AddStart uint64

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().Uint64Var(&AddEach, "each", 0, "period; number of intervals after which the same amount is added (0 if you do not want it to repeat)")
	addCmd.Flags().Uint64Var(&AddStart, "start", 0, "when do you want to add the amount (0 = at the very beginning of the plan)")
	//addCmd.Flags().BoolVar(&AddExclude, "exclude-start", false, "set this flag if you want to add the amount AFTER the date")
}
