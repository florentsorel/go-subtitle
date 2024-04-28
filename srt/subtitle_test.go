package srt

import (
	"testing"
	"time"
)

func TestString(t *testing.T) {
	data := []struct {
		input    Subtitle
		expected string
	}{
		{
			Subtitle{Number: 1, Start: time.Duration(1000000000), End: time.Duration(5000000000), Text: "First subtitle!"},
			"Number: 1, Start: 1s, End: 5s, Text: First subtitle!",
		},
		{
			Subtitle{Number: 2, Start: time.Duration(1000000000), End: time.Duration(5000000000), Text: "First subtitle!\nNew line"},
			"Number: 2, Start: 1s, End: 5s, Text: First subtitle!\nNew line",
		},
	}

	for _, tc := range data {
		t.Run(tc.expected, func(t *testing.T) {
			got := tc.input.String()
			if tc.expected != got {
				t.Errorf("expected:\n%s\ngot:\n%s", tc.expected, got)
			}
		})
	}
}
