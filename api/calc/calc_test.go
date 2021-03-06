package calc

import (
	"testing"
)

func TestSum(t *testing.T) {
	got, _ := getResult(1, 2, sum)
	if got != 3 {
		t.Errorf("1 + 2 = %v; want 3", got)
	}
}

func TestSubtract(t *testing.T) {
	got, _ := getResult(15, 5, subtract)
	if got != 10 {
		t.Errorf("15 - 5 = %v; want 10", got)
	}
}

func TestMultiply(t *testing.T) {
	got, _ := getResult(200, 5, multiply)
	if got != 1000 {
		t.Errorf("200 * 5 = %v; want 1000", got)
	}
}

func TestDivide(t *testing.T) {
	got, _ := getResult(100, 4, divide)
	if got != 25 {
		t.Errorf("100 / 4 = %v; want 25", got)
	}
}

func TestUnsupportedOperator(t *testing.T) {
	operator := "abc123"
	_, err := getResult(1, 1, operator)
	if err == nil {
		t.Errorf("Operator: %v is unsupported; want error != nil", operator)
	}
}
