package finaplan

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	plan := Init(DefaultConfig(), 5)
	assert.NoError(t, plan.Add("300", 2, 0))
	expectedProjection := []string{"300", "300", "600", "600", "900"}
	for i, amount := range plan.Projection {
		if !amount.Equal(decimal.RequireFromString(expectedProjection[i])) {
			t.Errorf("Element №%d = %s does not match the expected value (%v)", i, amount, expectedProjection[i])
		}
	}
}

func TestAddOnce(t *testing.T) {
	plan := Init(DefaultConfig(), 6)
	assert.NoError(t, plan.Add("12.3", 0, 2))
	expectedProjection := []string{"0", "0", "12.3", "12.3", "12.3", "12.3"}
	for i, amount := range plan.Projection {
		if !amount.Equal(decimal.RequireFromString(expectedProjection[i])) {
			t.Errorf("Element №%d = %s does not match the expected value (%v)", i, amount, expectedProjection[i])
		}
	}
}

func TestAddWithInvalidAmount(t *testing.T) {
	plan := Init(DefaultConfig(), 6)
	assert.Error(t, plan.Add("12.3$", 0, 2))
}
