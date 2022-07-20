package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sprkweb/finaplan-cli/internal/parser"
	"github.com/sprkweb/finaplan-cli/pkg/finaplan"
	"strconv"
)

// investCmd represents the invest command
var investCmd = &cobra.Command{
	Use:   "invest",
	Short: "Add some interest rate on top of your capital",
	Long: `Add some interest rate on top of your capital regularly

Let's say you want to invest your 300$ of savings and you expect 10% return per year.
Calculation interval in the example is 6 months, that means your interest is 10% per 2 intervals (6 * 2 = 12 months = 1 year)
$ finaplan init --each 6 --months | finaplan add 300 | \
    finaplan invest 1.1 --interval 2
---
interval_type: months
interval_length: 6
---
300
314.6426544510455
330
346.1069198961501
363`,
	Run: parser.ModifyPlan(func(plan *finaplan.FinancialPlan, args []string) error {
		interest, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			panic(err)
		}

		err = plan.Invest(interest, InvestInterval, InvestStart, InvestCompound)
		if err != nil {
			panic(err)
		}
		return nil
	}),
}

var InvestInterval uint64
var InvestStart uint64
var InvestCompound bool

func init() {
	rootCmd.AddCommand(investCmd)
	investCmd.Flags().Uint64Var(&InvestInterval, "interval", 1, "period; number of intervals after which the interest is applied")
	investCmd.Flags().Uint64Var(&InvestStart, "start", 0, "when do you want to start the investments (0 = at the very beginning of the plan)")
	investCmd.Flags().BoolVar(&InvestCompound, "compound", true, "set this flag if the specified interest is compound (all the interest in reinvested)")
}
