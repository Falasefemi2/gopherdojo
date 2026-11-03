package main

import "testing"

func TestNewCounter(t *testing.T) {
	counter := newCounter(0)
	if got := counter(); got != 1 {
		t.Fatalf("first call = %d, want 1", got)
	}
	if got := counter(); got != 2 {
		t.Errorf("second call = %d, want 2", got)
	}
	if got := counter(); got != 3 {
		t.Errorf("third call = %d, want 3", got)
	}

	other := newCounter(10)
	if got := other(); got != 11 {
		t.Errorf("newCounter(10) first call = %d, want 11", got)
	}
}

func TestCountersAreIndependent(t *testing.T) {
	first := newCounter(0)
	second := newCounter(0)

	if got := first(); got != 1 {
		t.Fatalf("first counter = %d, want 1", got)
	}
	if got := second(); got != 1 {
		t.Fatalf("second counter = %d, want 1; counters share state", got)
	}
	if got := first(); got != 2 {
		t.Errorf("first counter after second advanced = %d, want 2", got)
	}
}

func TestNewAccumulator(t *testing.T) {
	add, value := newAccumulator()

	if got := value(); got != 0 {
		t.Errorf("initial value = %d, want 0", got)
	}
	if got := add(5); got != 5 {
		t.Errorf("add(5) = %d, want 5", got)
	}
	if got := add(3); got != 8 {
		t.Errorf("add(3) = %d, want 8", got)
	}
	if got := value(); got != 8 {
		t.Errorf("value() = %d, want 8", got)
	}
	if got := value(); got != 8 {
		t.Errorf("second value() = %d, want 8; value must not change the total", got)
	}
	if got := add(-10); got != -2 {
		t.Errorf("add(-10) = %d, want -2", got)
	}
}

func TestAccumulatorsAreIndependent(t *testing.T) {
	addFirst, valueFirst := newAccumulator()
	addSecond, valueSecond := newAccumulator()

	addFirst(4)
	if got := valueSecond(); got != 0 {
		t.Errorf("second accumulator value = %d, want 0", got)
	}
	if got := addSecond(7); got != 7 {
		t.Errorf("second add(7) = %d, want 7", got)
	}
	if got := valueFirst(); got != 4 {
		t.Errorf("first accumulator value = %d, want 4", got)
	}
}

func TestNewMultiplier(t *testing.T) {
	triple := newMultiplier(3)
	if got := triple(4); got != 12 {
		t.Errorf("triple(4) = %d, want 12", got)
	}
	if got := triple(0); got != 0 {
		t.Errorf("triple(0) = %d, want 0", got)
	}

	factor := 2
	double := newMultiplier(factor)
	factor = 100
	if got := double(5); got != 10 {
		t.Errorf("double(5) = %d, want 10; the closure captured factor by value", got)
	}
}
