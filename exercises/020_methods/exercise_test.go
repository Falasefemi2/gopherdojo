package main

import "testing"

func TestCounterMutation(t *testing.T) {
	var c Counter
	c.Inc()
	c.Inc()
	c.Add(5)
	c.Add(-2)

	if got, want := c.Value(), 5; got != want {
		t.Errorf("after Inc, Inc, Add(5), Add(-2) Value() = %d, want %d", got, want)
	}
}

func TestPointerReceiverSharesState(t *testing.T) {
	var c Counter
	p := &c
	q := &c

	p.Inc()
	q.Add(4)

	if got, want := c.Value(), 5; got != want {
		t.Errorf("shared counter Value() = %d, want %d", got, want)
	}
}

func TestZeroCounter(t *testing.T) {
	var c Counter
	if got := c.Value(); got != 0 {
		t.Errorf("zero Counter Value() = %d, want 0", got)
	}
}

func TestWithAdded(t *testing.T) {
	tests := []struct {
		name  string
		start int
		delta int
		want  int
	}{
		{"positive delta", 3, 4, 7},
		{"zero delta", 3, 0, 3},
		{"negative delta", 3, -5, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Counter{n: tt.start}
			got := c.WithAdded(tt.delta)

			if got.Value() != tt.want {
				t.Errorf("Counter{%d}.WithAdded(%d).Value() = %d, want %d", tt.start, tt.delta, got.Value(), tt.want)
			}
			if c.Value() != tt.start {
				t.Errorf("WithAdded mutated the receiver: Value() = %d, want %d", c.Value(), tt.start)
			}
		})
	}

	t.Run("result is independent", func(t *testing.T) {
		c := Counter{n: 3}
		snapshot := c.WithAdded(4)
		snapshot.Inc()
		if got := c.Value(); got != 3 {
			t.Errorf("mutating the returned copy changed the original to %d, want 3", got)
		}
	})
}
