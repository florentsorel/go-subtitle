package srt

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	testCases := []struct {
		input          string
		expected       time.Duration
		expectedString string
	}{
		{"00:01:20,123", time.Duration(0*time.Hour + 1*time.Minute + 20*time.Second + 123*time.Millisecond), "1m20.123s"},
		{"02:14:22,456", time.Duration(2*time.Hour + 14*time.Minute + 22*time.Second + 456*time.Millisecond), "2h14m22.456s"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, err := parseTime(tc.input)
			if err != nil {
				t.Error(err)
			}

			if result != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, result)
			}

			if result.String() != tc.expectedString {
				t.Errorf("expected %s, got %s", tc.expectedString, result.String())
			}
		})
	}
}
