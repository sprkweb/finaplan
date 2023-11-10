package cmd

import (
	"github.com/sprkweb/finaplan/internal/parser"

	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the given plan in a more human-friendly manner",
	Long: `Print the given plan in a more human-friendly manner

Example:
$ finaplan init --intervals 5 --months | \
    finaplan add 10000 --each 1 | \
        finaplan print
Month 0         | 10000.00
Month 1         | 20000.00
Month 2         | 30000.00
Month 3         | 40000.00
Month 4         | 50000.00
`,
	Run: func(cmd *cobra.Command, args []string) {
		plan, err := parser.ParsePlanFromStdin()
		if err != nil {
			panic(err)
		}

		if err := parser.PrettyPrintPlanToStdout(plan); err != nil {
			panic(err)
		}
	},
}

var PrintGraph bool
var PrintNoColor bool

func init() {
	rootCmd.AddCommand(printCmd)
}
