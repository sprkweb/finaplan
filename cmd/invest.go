package cmd

import (
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"github.com/sprkweb/finaplan-cli/finaplan/internal/parser"
	"github.com/sprkweb/finaplan-cli/finaplan/pkg/finaplan"
)

// investCmd represents the invest command
var investCmd = &cobra.Command{
	Use:   "invest INTEREST",
	Short: "Add some interest rate on top of your capital",
	Long: `Add some interest rate on top of your capital regularly

Let's say you want to invest your 300$ of savings and you expect 10% return per year.
Calculation interval in the example is 6 months, that means your interest is 10% per 2 intervals (6 * 2 = 12 months = 1 year)
$ finaplan init --each 6 --months | finaplan add 300 | \
    finaplan invest 0.1 --interval 2
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
		interest, err := decimal.NewFromString(args[0])
		if err != nil {
			return err
		}

		err = plan.Invest(interest, InvestInterval, InvestStart, !InvestSimple)
		if err != nil {
			return err
		}
		return nil
	}),
}

var InvestInterval uint32
var InvestStart uint32
var InvestSimple bool

func init() {
	rootCmd.AddCommand(investCmd)
	investCmd.Flags().Uint32Var(&InvestInterval, "interval", 1, "period; number of intervals after which the interest is applied")
	investCmd.Flags().Uint32Var(&InvestStart, "start", 0, "when do you want to start the investments (0 = at the very beginning of the plan)")
	investCmd.Flags().BoolVar(&InvestSimple, "simple", false, "set this flag if the specified interest is simple (= not compound, the interest is not reinvested)")
}
