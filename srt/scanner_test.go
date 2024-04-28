package srt

import (
	"errors"
	"testing"
	"time"
)

func TestScan(t *testing.T) {
	data := []struct {
		input    string
		expected *Subtitle
		err      error
	}{
		{
			`1
00:00:01,000 --> 00:00:05,000
First subtitle!`,
			&Subtitle{Number: 1, Start: time.Duration(1000000000), End: time.Duration(5000000000), Text: "First subtitle!"},
			nil,
		},
		{
			`2
00:00:01,000 --> 00:00:05,000
First subtitle!
New line`,
			&Subtitle{Number: 2, Start: time.Duration(1000000000), End: time.Duration(5000000000), Text: "First subtitle!\nNew line"},
			nil,
		},
		{
			`
00:00:01,000 --> 00:00:05,000
Empty number`,
			nil,
			errors.New("empty number for:\n\"\n00:00:01,000 --> 00:00:05,000\nEmpty number\""),
		},
		{
			`00:00:01,000 --> 00:00:05,000
Invalid number`,
			nil,
			errors.New("invalid number: \"00:00:01,000 --> 00:00:05,000\""),
		},
		{
			`one
00:00:01,000 --> 00:00:05,000
Invalid number, cannot convert string to int`,
			nil,
			errors.New("invalid number: \"one\""),
		},
		{
			`3
00:00:01,000 -> 00:00:05,000
Invalid time format!`,
			nil,
			errors.New("invalid time format: \"00:00:01,000 -> 00:00:05,000\""),
		},
	}

	for _, tc := range data {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parse(tc.input)
			if err != nil {
				if tc.err.Error() != err.Error() {
					t.Errorf("expected: %s, got: %s", tc.err, err)
				}
			} else {
				if tc.expected.String() != got.String() {
					t.Errorf("expected:\n%s\ngot:\n%s", tc.expected, got)
				}
			}
		})
	}
}
