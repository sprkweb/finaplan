package parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/shopspring/decimal"
	"github.com/sprkweb/finaplan/pkg/finaplan"
	"golang.org/x/term"
)

func PrettyPrintPlanToStdout(plan *finaplan.FinancialPlan, graph bool, colored bool) error {
	if len(plan.Projection) < 1 {
		return nil
	}

	unit, err := finaplan.GetIntervalUnit(plan.Config.IntervalType)
	if err != nil {
		return err
	}
	// capitalize
	unit = fmt.Sprintf("%s%s", strings.ToUpper(unit[0:1]), unit[1:])

	labelLength := maxLabelLength(plan, unit)

	var format string
	var symbolValue decimal.Decimal
	if graph {
		numberColWidth := maxPrintLen(plan.Projection)
		graphWidth := getTerminalWidth() - labelLength - numberColWidth - 2
		symbolValue = maxVal(plan).Div(decimal.NewFromInt(int64(graphWidth)))
		format = fmt.Sprintf("%%%ds %%s %%v\n", labelLength)
	} else {
		format = fmt.Sprintf("%%%ds | %%v\n", labelLength)
	}

	for i, v := range plan.Projection {
		interval := uint64(plan.Config.IntervalLength) * uint64(i)
		label := sprintLabel(unit, interval)

		if graph {
			barsNum := int(v.Div(symbolValue).Floor().IntPart())
			barStr, err := sprintBar(barsNum, colored)
			if err != nil {
				return err
			}
			fmt.Printf(format, label, barStr, sprintNum(v))
		} else {
			fmt.Printf(format, label, sprintNum(v))
		}
	}

	return nil
}

func sprintLabel(unit string, interval uint64) string {
	return fmt.Sprintf("%s %d", unit, interval)
}

func sprintBar(barsNum int, colored bool) (string, error) {
	var bar strings.Builder
	for i := 0; i < barsNum && i < 200; i++ {
		_, err := bar.WriteRune('$')
		if err != nil {
			return "", err
		}
	}
	barStr := bar.String()
	if colored {
		greenColor := color.New(color.BgGreen).Add(color.FgBlack)
		barStr = greenColor.Sprint(barStr)
	}
	return barStr, nil
}

func sprintNum(num decimal.Decimal) string {
	return num.StringFixedBank(2)
}

func maxVal(plan *finaplan.FinancialPlan) decimal.Decimal {
	return decimal.Max(plan.Projection[0], plan.Projection[1:]...)
}

func maxPrintLen(arr finaplan.Projection) int {
	if len(arr) < 1 {
		return 0
	}
	length := func(x decimal.Decimal) int {
		return len(sprintNum(x))
	}
	max := length(arr[0])
	for _, v := range arr {
		l := length(v)
		if l > max {
			max = l
		}
	}
	return max
}

func maxLabelLength(plan *finaplan.FinancialPlan, unit string) int {
	return len(sprintLabel(unit, uint64(len(plan.Projection)-1)))
}

const defaultWidth = 50

func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if width <= 0 || err != nil {
		return defaultWidth
	} else {
		return width
	}
}
