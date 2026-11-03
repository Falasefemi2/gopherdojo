package main

import (
	"reflect"
	"testing"
)

func TestFormatAll(t *testing.T) {
	tests := []struct {
		name  string
		label string
		p     Point
		n     int
		f     float64
		s     string
		want  Formatted
	}{
		{
			name:  "origin",
			label: "origin",
			p:     Point{X: 0, Y: 0},
			n:     0,
			f:     0,
			s:     "",
			want: Formatted{
				Value:    "{0 0}",
				Plus:     "{X:0 Y:0}",
				GoSyntax: "main.Point{X:0, Y:0}",
				Type:     "main.Point",
				Int:      "0",
				Float:    "0.00",
				Quoted:   `""`,
				Label:    "origin",
			},
		},
		{
			name:  "typical values",
			label: "checkout",
			p:     Point{X: 3, Y: 4},
			n:     42,
			f:     3.14159,
			s:     "ledger",
			want: Formatted{
				Value:    "{3 4}",
				Plus:     "{X:3 Y:4}",
				GoSyntax: "main.Point{X:3, Y:4}",
				Type:     "main.Point",
				Int:      "42",
				Float:    "3.14",
				Quoted:   `"ledger"`,
				Label:    "checkout",
			},
		},
		{
			name:  "negative and escaped",
			label: "audit",
			p:     Point{X: -1, Y: 2},
			n:     -7,
			f:     2.5,
			s:     `path "x"`,
			want: Formatted{
				Value:    "{-1 2}",
				Plus:     "{X:-1 Y:2}",
				GoSyntax: "main.Point{X:-1, Y:2}",
				Type:     "main.Point",
				Int:      "-7",
				Float:    "2.50",
				Quoted:   `"path \"x\""`,
				Label:    "audit",
			},
		},
		{
			name:  "rounding carries",
			label: "price",
			p:     Point{X: 9, Y: 9},
			n:     1,
			f:     9.999,
			s:     "a\tb",
			want: Formatted{
				Value:    "{9 9}",
				Plus:     "{X:9 Y:9}",
				GoSyntax: "main.Point{X:9, Y:9}",
				Type:     "main.Point",
				Int:      "1",
				Float:    "10.00",
				Quoted:   `"a\tb"`,
				Label:    "price",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatAll(tt.label, tt.p, tt.n, tt.f, tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatAll(%q, %+v, %d, %v, %q) =\n  %+v\nwant\n  %+v",
					tt.label, tt.p, tt.n, tt.f, tt.s, got, tt.want)
			}
		})
	}
}
