package finaplan

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestSimpleInvest(t *testing.T) {
	plan := Init(DefaultConfig(), 7)
	addAmount := decimal.NewFromInt(300)
	interest := decimal.RequireFromString("0.1") // 10%
	plan.Add(addAmount, 0, 0)
	if err := plan.Invest(interest, 2, 2, false); err != nil {
		t.Errorf("Got error: %s", err)
	}

	expectedProjection := []string{"300", "300", "300", "315", "330", "345", "360"}
	for i, amount := range plan.Projection {
		if !amount.Equal(decimal.RequireFromString(expectedProjection[i])) {
			t.Errorf("Element №%d = %v does not match the expected value (%v)", i, amount, expectedProjection)
		}
	}
}

func TestInvestWithAdd(t *testing.T) {
	plan := Init(DefaultConfig(), 6)
	addAmount := decimal.NewFromInt(300)
	interest := decimal.RequireFromString("0.1") // 10%
	plan.Add(addAmount, 1, 0)
	if err := plan.Invest(interest, 2, 2, false); err != nil {
		t.Errorf("Got error: %s", err)
	}

	expectedProjection := []string{"300", "600", "900", "1245", "1605", "1980"}
	for i, amount := range plan.Projection {
		if !amount.Equal(decimal.RequireFromString(expectedProjection[i])) {
			t.Errorf("Element №%d = %v does not match the expected value (%v)", i, amount, expectedProjection[i])
		}
	}
}

func TestInvestWithOneNumber(t *testing.T) {
	plan := Init(DefaultConfig(), 1)
	addAmount := decimal.NewFromInt(300)
	interest := decimal.RequireFromString("0.1") // 10%
	plan.Add(addAmount, 1, 0)
	if err := plan.Invest(interest, 2, 2, false); err != nil {
		t.Errorf("Got error: %s", err)
	}

	if !plan.Projection[0].Equal(addAmount) {
		t.Errorf("Element %v does not match the expected value %v", plan.Projection[0], 300)
	}
}

func TestInvestWithNoNumbers(t *testing.T) {
	plan := Init(DefaultConfig(), 0)
	addAmount := decimal.NewFromInt(300)
	interest := decimal.RequireFromString("0.1") // 10%
	plan.Add(addAmount, 1, 0)
	if err := plan.Invest(interest, 2, 2, false); err != nil {
		t.Errorf("Got error: %s", err)
	}
}

func TestInvestCompound(t *testing.T) {
	plan := Init(DefaultConfig(), 5)
	addAmount := decimal.NewFromInt(300)
	interest := decimal.RequireFromString("0.21") // 21%
	plan.Add(addAmount, 0, 0)
	if err := plan.Invest(interest, 2, 0, true); err != nil {
		t.Errorf("Got error: %s", err)
	}

	expectedProjection := []string{"300", "330", "363", "399.3", "439.23"}
	for i, amount := range plan.Projection {
		if !amount.Equal(decimal.RequireFromString(expectedProjection[i])) {
			t.Errorf("Element №%d = %v does not match the expected value (%v)", i, amount, expectedProjection[i])
		}
	}
}

func TestInvestCompoundWithAdd(t *testing.T) {
	plan := Init(DefaultConfig(), 9)
	addAmount := decimal.NewFromInt(200)
	interest := decimal.RequireFromString("0.1") // 10%
	plan.Add(addAmount, 3, 2)
	if err := plan.Invest(interest, 1, 1, true); err != nil {
		t.Errorf("Got error: %s", err)
	}

	expectedProjection := []string{"0", "0", "200", "220", "242", "466.2", "512.82", "564.102", "820.5122"}
	for i, amount := range plan.Projection {
		if !amount.Equal(decimal.RequireFromString(expectedProjection[i])) {
			t.Errorf("Element №%d = %v does not match the expected value (%v)", i, amount, expectedProjection[i])
		}
	}
}
