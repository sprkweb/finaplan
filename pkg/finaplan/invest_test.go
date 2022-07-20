package finaplan

import "testing"

const eps = 0.00001

func equal(a, b float64) bool {
	return a < (b+eps) && a > (b-eps)
}

func TestSimpleInvest(t *testing.T) {
	plan := Init(DefaultConfig(), 7)
	plan.Add(300, 0, 0)
	if err := plan.Invest(1.1, 2, 2, false); err != nil {
		t.Errorf("Got error: %s", err)
	}

	expectedProjection := []float64{300, 300, 300, 315, 330, 345, 360}
	for i, amount := range plan.Projection {
		if !equal(float64(amount), expectedProjection[i]) {
			t.Errorf("Element №%d = %v does not match the expected value (%v)", i, amount, expectedProjection)
		}
	}
}

func TestInvestWithAdd(t *testing.T) {
	plan := Init(DefaultConfig(), 6)
	plan.Add(300, 1, 0)
	if err := plan.Invest(1.1, 2, 2, false); err != nil {
		t.Errorf("Got error: %s", err)
	}

	expectedProjection := []float64{300, 600, 900, 1245, 1605, 1980}
	for i, amount := range plan.Projection {
		if !equal(float64(amount), expectedProjection[i]) {
			t.Errorf("Element №%d = %v does not match the expected value (%v)", i, amount, expectedProjection[i])
		}
	}
}

func TestInvestWithOneNumber(t *testing.T) {
	plan := Init(DefaultConfig(), 1)
	plan.Add(300, 1, 0)
	if err := plan.Invest(1.1, 2, 2, false); err != nil {
		t.Errorf("Got error: %s", err)
	}

	if !equal(float64(plan.Projection[0]), 300) {
		t.Errorf("Element %v does not match the expected value %v", plan.Projection[0], 300)
	}
}

func TestInvestWithNoNumbers(t *testing.T) {
	plan := Init(DefaultConfig(), 0)
	plan.Add(300, 1, 0)
	if err := plan.Invest(1.1, 2, 2, false); err != nil {
		t.Errorf("Got error: %s", err)
	}
}

func TestInvestCompound(t *testing.T) {
	plan := Init(DefaultConfig(), 5)
	plan.Add(300, 0, 0)
	if err := plan.Invest(1.21, 2, 0, true); err != nil {
		t.Errorf("Got error: %s", err)
	}

	expectedProjection := []float64{300, 330, 363, 399.3, 439.23}
	for i, amount := range plan.Projection {
		if !equal(float64(amount), expectedProjection[i]) {
			t.Errorf("Element №%d = %v does not match the expected value (%v)", i, amount, expectedProjection[i])
		}
	}
}

func TestInvestCompoundWithAdd(t *testing.T) {
	plan := Init(DefaultConfig(), 9)
	plan.Add(200, 3, 2)
	if err := plan.Invest(1.1, 1, 1, true); err != nil {
		t.Errorf("Got error: %s", err)
	}

	expectedProjection := []float64{0, 0, 200, 220, 242, 466.2, 512.82, 564.102, 820.5122}
	for i, amount := range plan.Projection {
		if !equal(float64(amount), expectedProjection[i]) {
			t.Errorf("Element №%d = %v does not match the expected value (%v)", i, amount, expectedProjection[i])
		}
	}
}
