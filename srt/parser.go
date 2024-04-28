package srt

import (
	"fmt"
	"regexp"
	"time"
)

func parseTime(value string) (time.Duration, error) {
	regex := regexp.MustCompile(`(\d{2}):(\d{2}):(\d{2}),(\d{3})`)
	matches := regex.FindStringSubmatch(value)

	if len(matches) != 5 {
		return 0, fmt.Errorf("invalid time format: %s", value)
	}

	hour, err := time.ParseDuration(matches[1] + "h")
	if err != nil {
		return 0, err
	}

	minute, err := time.ParseDuration(matches[2] + "m")
	if err != nil {
		return 0, err
	}

	second, err := time.ParseDuration(matches[3] + "s")
	if err != nil {
		return 0, err
	}

	millisecond, err := time.ParseDuration(matches[4] + "ms")
	if err != nil {
		return 0, err
	}

	return time.Duration(hour + minute + second + millisecond), nil
}
