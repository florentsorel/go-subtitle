package srt

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	data := []struct {
		input         string
		subtitleCount int
		err           error
	}{
		{
			"./../testdata/srt/five_subtitle.srt",
			5,
			nil,
		},
		{
			"./../testdata/srt/invalid_separator.srt",
			0,
			errors.New("invalid time format: \"00:01:10,012 -> 00:01:13,001\""),
		},
	}

	for _, tc := range data {
		t.Run(filepath.Clean(tc.input), func(t *testing.T) {
			got, err := Open(tc.input)
			if err != nil {
				if tc.err.Error() != err.Error() {
					t.Errorf("expected: %s, got: %s", tc.err, err)
				}
			} else {
				if tc.subtitleCount != len(got) {
					t.Errorf("expected: %d, got: %d", tc.subtitleCount, len(got))
				}
			}
		})
	}
}
