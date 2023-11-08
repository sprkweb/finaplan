package finaplan

import "github.com/shopspring/decimal"

func percent(paramName string, n string) (decimal.Decimal, error) {
	if n[len(n)-1] != byte('%') {
		d, err := decimal.NewFromString(n)
		if err != nil {
			return decimal.Decimal{}, NewErrDecodeDecimal(paramName, err)
		}
		return d, nil
	}

	d, err := decimal.NewFromString(n[:len(n)-1])
	if err != nil {
		return decimal.Decimal{}, NewErrDecodeDecimal(paramName, err)
	}
	return d.Shift(-2), nil
}
