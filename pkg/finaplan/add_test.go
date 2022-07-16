package finaplan

import (
	"testing"
)

func TestAdd(t *testing.T) {
	plan := Init(DefaultConfig(), 5)
	plan.Add(300, 2, 0)
	expectedProjection := []float64{300, 300, 600, 600, 900}
	for i, amount := range expectedProjection {
		if amount != expectedProjection[i] {
			t.Errorf("ELement №%d = %v does not match the expected value (%v)", i, amount, expectedProjection)
		}
	}
}

func TestAddOnce(t *testing.T) {
	plan := Init(DefaultConfig(), 6)
	plan.Add(12.3, 0, 2)
	expectedProjection := []float64{0, 0, 12.3, 12.3, 12.3, 12.3}
	for i, amount := range expectedProjection {
		if amount != expectedProjection[i] {
			t.Errorf("ELement №%d = %v does not match the expected value (%v)", i, amount, expectedProjection)
		}
	}
}
