package cmd

import (
	"github.com/sprkweb/finaplan-cli/internal/parser"

	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the given plan in a more human-friendly manner",
	Long: `Print the given plan in a more human-friendly manner

Example:
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

func init() {
	rootCmd.AddCommand(printCmd)
}
