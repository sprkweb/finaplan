package parser

import "github.com/shopspring/decimal"

func Percent(input string) (decimal.Decimal, error) {
	if input[len(input)-1] != byte('%') {
		return decimal.NewFromString(input)
	}

	percent, err := decimal.NewFromString(input[:len(input)-1])
	if err != nil {
		return decimal.Decimal{}, err
	}

	return percent.Shift(-2), nil
}
