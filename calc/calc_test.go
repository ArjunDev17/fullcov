package calc

import (
	"testing"
)

func TestAddSubMul(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add: got %d want %d", got, 5)
	}
	if got := Sub(10, 4); got != 6 {
		t.Fatalf("Sub: got %d want %d", got, 6)
	}
	if got := Mul(3, 7); got != 21 {
		t.Fatalf("Mul: got %d want %d", got, 21)
	}
}

func TestDiv(t *testing.T) {
	// normal division
	res, err := Div(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != 5 {
		t.Fatalf("Div: got %v want %v", res, 5)
	}

	// division by zero branch
	_, err = Div(1, 0)
	if err == nil {
		t.Fatalf("expected error dividing by zero")
	}
}

func TestFactorial(t *testing.T) {
	// 0! = 1
	if v, err := Factorial(0); err != nil || v != 1 {
		t.Fatalf("Factorial 0: got %v, %v; want 1, nil", v, err)
	}

	// 5! = 120
	if v, err := Factorial(5); err != nil || v != 120 {
		t.Fatalf("Factorial 5: got %v, %v; want 120, nil", v, err)
	}

	// negative input
	if _, err := Factorial(-3); err == nil {
		t.Fatalf("Factorial negative: expected error")
	}
}

func TestReverse(t *testing.T) {
	// ascii
	if got := Reverse("hello"); got != "olleh" {
		t.Fatalf("Reverse ascii: got %v want %v", got, "olleh")
	}
	// utf8
	if got := Reverse("世世🌏"); got != "🌏世世" {
		t.Fatalf("Reverse utf8: got %v want %v", got, "🌏世世")
	}
}

func TestIsPrime(t *testing.T) {
	cases := map[int]bool{
		-1: false,
		0:  false,
		1:  false,
		2:  true,
		3:  true,
		4:  false,
		17: true,
		18: false,
		97: true,
	}
	for n, want := range cases {
		if got := IsPrime(n); got != want {
			t.Fatalf("IsPrime(%d): got %v want %v", n, got, want)
		}
	}
}
