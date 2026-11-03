package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestLogInfo(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want string
	}{
		{"plain", "worker started", "INFO: worker started\n"},
		{"empty", "", "INFO: \n"},
		{"already newline", "done\n", "INFO: done\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logInfo(&buf, tt.msg)

			got := buf.String()
			if got != tt.want {
				t.Errorf("logInfo(%q) wrote %q, want %q", tt.msg, got, tt.want)
			}
			if strings.Count(got, "\n") != 1 {
				t.Errorf("logInfo(%q) wrote %d newlines, want exactly 1", tt.msg, strings.Count(got, "\n"))
			}
		})
	}
}

func TestLogError(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want string
	}{
		{"plain", "queue empty", "ERROR: queue empty\n"},
		{"already newline", "timeout\n", "ERROR: timeout\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logError(&buf, tt.msg)

			got := buf.String()
			if got != tt.want {
				t.Errorf("logError(%q) wrote %q, want %q", tt.msg, got, tt.want)
			}
			if strings.Contains(got, "202") {
				t.Errorf("logError(%q) wrote %q, want no timestamp", tt.msg, got)
			}
		})
	}
}
