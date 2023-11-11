package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sprkweb/finaplan/internal/finaplanio"
	"github.com/sprkweb/finaplan/pkg/finaplan"
)

// inflationCmd represents the inflation command
var inflationCmd = &cobra.Command{
	Use:   "inflation INFLATION",
	Short: "Adjust all the previous plan for inflation",
	Long: `Adjust all the previous plan for some expected inflation

Let's see how much your 300$ dollars would worth in 2 years with 4% inflation.
In this example we represent a year as 2 intervals of 6 months.

$ finaplan init --each 6 --months --intervals 5 | finaplan add 300 | \
    finaplan inflation 4% --interval 2 | finaplan print
Month 0         | 300.00
Month 6         | 294.17
Month 12        | 288.46
Month 18        | 282.86
Month 24        | 277.37
`,
	Run: finaplanio.ModifyPlan(func(plan *finaplan.FinancialPlan, args []string) error {
		return plan.Inflation(args[0], InflationInterval)
	}),
}

var InflationInterval uint32

func init() {
	rootCmd.AddCommand(inflationCmd)
	inflationCmd.Flags().Uint32Var(&InflationInterval, "interval", 1, "period; number of intervals after which the interest is applied")
}
