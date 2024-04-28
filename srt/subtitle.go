package srt

import (
	"fmt"
	"time"
)

type Subtitle struct {
	Number int
	Start  time.Duration
	End    time.Duration
	Text   string
}

func (s *Subtitle) String() string {
	return fmt.Sprintf("Number: %d, Start: %s, End: %s, Text: %s", s.Number, s.Start, s.End, s.Text)
}
