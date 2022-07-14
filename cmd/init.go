package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize an empty plan",
	Long: `Initialize an empty financial plan

Example:
Initialize an empty plan for 15 weeks ahead with the calculation interval of 3 weeks:
$ finaplan init --each 3 --weeks --intervals 5
---
interval_type: weeks
interval_length: 3
interval_amount: 5 
---
0
0
0
0
0
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
