package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(2, 2)
	if total != 4 {
		t.Errorf("Sum was incorrect, Actual: %d, Expected: %d", total, 4)
	}
	t.Log("running TestAdd")
}

func TestSub(t *testing.T) {
	total := Subtract(2, 2)
	if total != 0 {
		t.Errorf("Sum was incorrect, Actual: %d, Expected: %d", total, 0)
	}
	t.Log("running TestSubtract")
}
