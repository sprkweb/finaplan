package finaplan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	plan := Init(DefaultConfig(), 5)
	assert.NoError(t, plan.Add("300", 2, 0))
	want := []string{"300", "300", "600", "600", "900"}
	assert.Equal(t, want, plan.Print())
}

func TestAddOnce(t *testing.T) {
	plan := Init(DefaultConfig(), 6)
	assert.NoError(t, plan.Add("12.3", 0, 2))
	want := []string{"0", "0", "12.3", "12.3", "12.3", "12.3"}
	assert.Equal(t, want, plan.Print())
}

func TestAddWithInvalidAmount(t *testing.T) {
	plan := Init(DefaultConfig(), 6)
	assert.Error(t, plan.Add("12.3$", 0, 2))
}
