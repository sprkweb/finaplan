package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sprkweb/finaplan-cli/internal/parser"
	"github.com/sprkweb/finaplan-cli/pkg/finaplan"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize an empty plan",
	Long: `Initialize an empty financial plan

Example:
Initialize an empty plan for 15 weeks ahead
with the calculation interval of 3 weeks:
$ finaplan init --each 3 --weeks --intervals 5
---
interval_type: weeks
interval_length: 3
---
0
0
0
0
0
`,
	Run: func(cmd *cobra.Command, args []string) {
		newPlan := finaplan.Init(&finaplan.PlanConfig{
			IntervalType:   IntervalType(),
			IntervalLength: IntervalLength,
		}, IntervalAmount)
		planStr, err := parser.PrintPlan(newPlan)
		if err != nil {
			panic(err)
		}
		fmt.Print(planStr)
	},
}

var IntervalLength uint32
var IntervalAmount uint64
var IntervalType func() finaplan.IntervalType

func init() {
	rootCmd.AddCommand(initCmd)

	defaultConfig := finaplan.DefaultConfig()
	initCmd.Flags().Uint32Var(&IntervalLength, "each", defaultConfig.IntervalLength, "amount of given units in an interval")
	initCmd.Flags().Uint64Var(&IntervalAmount, "intervals", 5, "amount of defined intervals to calculate")

	getUnitFlag := func(unit string) *bool {
		description := fmt.Sprintf("set this argument to calculate the plan in %s", unit)
		return initCmd.Flags().Bool(unit, false, description)
	}

	isDays := getUnitFlag("days")
	isWeeks := getUnitFlag("weeks")
	isMonths := getUnitFlag("months")
	isYears := getUnitFlag("years")

	IntervalType = func() finaplan.IntervalType {
		switch {
		case *isDays:
			return finaplan.Days
		case *isWeeks:
			return finaplan.Weeks
		case *isMonths:
			return finaplan.Months
		case *isYears:
			return finaplan.Years
		default:
			return defaultConfig.IntervalType
		}
	}
}
