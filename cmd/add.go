package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sprkweb/finaplan-cli/internal/parser"
	"github.com/sprkweb/finaplan-cli/pkg/finaplan"
	"io"
	"os"
	"strconv"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a certain amount of money to your plan",
	Long: `Add a certain amount of money to your financial plan (e.g. savings)

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
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			panic(err)
		}

		reader := bufio.NewReader(os.Stdin)
		input, err := io.ReadAll(reader)
		if err != nil {
			panic(err)
		}

		plan, err := parser.ParsePlan(string(input))
		if err != nil {
			panic(err)
		}

		plan.Add(finaplan.ProjectionUnit(amount), Each, Start)
		planStr, err := parser.PrintPlan(plan)
		if err != nil {
			panic(err)
		}
		fmt.Println(planStr)
	},
}

var Each uint64
var Start uint64

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().Uint64Var(&Each, "each", 0, "period; number of intervals after which the same amount is added (0 if you do not want it to repeat)")
	addCmd.Flags().Uint64Var(&Start, "start", 0, "when do you want to add the amount (0 = at the very beginning of the plan)")
	//addCmd.Flags().BoolVar(&Exclude, "exclude-start", false, "set this flag if you want to add the amount AFTER the date")
}
