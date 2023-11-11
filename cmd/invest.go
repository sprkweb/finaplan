package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sprkweb/finaplan/internal/finaplanio"
	"github.com/sprkweb/finaplan/pkg/finaplan"
)

// investCmd represents the invest command
var investCmd = &cobra.Command{
	Use:   "invest INTEREST",
	Short: "Add some interest rate on top of your capital",
	Long: `Add some interest rate on top of your capital regularly

Let's say you want to invest your 300$ of savings and you expect 10% return per year.
Calculation interval in the example is 6 months, that means your interest is 10% per 2 intervals (6 * 2 = 12 months = 1 year)
$ finaplan init --each 6 --months --intervals 5 | finaplan add 300 | \
    finaplan invest 10% --interval 2 | finaplan print
Month 0         | 300.00
Month 6         | 314.64
Month 12        | 330.00
Month 18        | 346.11
Month 24        | 363.00
`,
	Run: finaplanio.ModifyPlan(func(plan *finaplan.FinancialPlan, args []string) error {
		return plan.Invest(args[0], InvestInterval, InvestStart, !InvestSimple)
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
