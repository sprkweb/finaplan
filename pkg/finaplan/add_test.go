package finaplan

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestAdd(t *testing.T) {
	plan := Init(DefaultConfig(), 5)
	addAmount := decimal.NewFromInt(300)
	plan.Add(addAmount, 2, 0)
	expectedProjection := []string{"300", "300", "600", "600", "900"}
	for i, amount := range plan.Projection {
		if !amount.Equal(decimal.RequireFromString(expectedProjection[i])) {
			t.Errorf("Element №%d = %s does not match the expected value (%v)", i, amount, expectedProjection[i])
		}
	}
}

func TestAddOnce(t *testing.T) {
	plan := Init(DefaultConfig(), 6)
	addAmount := decimal.RequireFromString("12.3")
	plan.Add(addAmount, 0, 2)
	expectedProjection := []string{"0", "0", "12.3", "12.3", "12.3", "12.3"}
	for i, amount := range plan.Projection {
		if !amount.Equal(decimal.RequireFromString(expectedProjection[i])) {
			t.Errorf("Element №%d = %s does not match the expected value (%v)", i, amount, expectedProjection[i])
		}
	}
}
