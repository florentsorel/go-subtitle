package srt

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const durationSeparator = " --> "

type scanner struct {
	scanner bufio.Scanner
	next    Subtitle
	err     error
}

func newScanner(r io.Reader) *scanner {
	sc := bufio.NewScanner(r)
	sc.Split(scan)
	return &scanner{scanner: *sc, next: Subtitle{}, err: nil}
}

func (s *scanner) Scan() (wasRead bool) {
	if s.scanner.Scan() {
		subtitle, err := parse(s.scanner.Text())
		if err != nil {
			s.err = err
			return false
		}

		s.next = *subtitle
		return true
	}
	return false
}

func (s *scanner) Subtitle() Subtitle {
	return s.next
}

func (s *scanner) Err() error {
	if s.err != nil {
		return s.err
	}
	return s.scanner.Err()
}

func parse(line string) (*Subtitle, error) {
	subtitle := new(Subtitle)

	elements := strings.Split(line, "\n")
	for i := 0; i < len(elements); i++ {
		text := strings.TrimRight(elements[i], "\r")
		switch i {
		case 0:
			if text == "" {
				return nil, fmt.Errorf("empty number for:\n\"%v\"", line)
			}

			nextNumber, err := strconv.Atoi(text)
			if err != nil {
				return nil, fmt.Errorf("invalid number: \"%v\"", text)
			}

			subtitle.Number = nextNumber
		case 1:
			times := strings.Split(text, durationSeparator)
			if len(times) != 2 {
				return nil, fmt.Errorf("invalid time format: \"%v\"", text)
			}

			startTime, err := parseTime(times[0])
			if err != nil {
				return nil, fmt.Errorf("invalid start time: \"%v\"", times[0])
			}

			endTime, err := parseTime(times[1])
			if err != nil {
				return nil, fmt.Errorf("invalid start time: \"%v\"", times[1])
			}

			subtitle.Start = startTime
			subtitle.End = endTime
		default:
			if len(subtitle.Text) > 0 {
				subtitle.Text += "\n"
			}
			subtitle.Text += text
		}
	}

	return subtitle, nil
}

func scan(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\n', '\r', '\n'}); i >= 0 {
		return i + 3, dropCR(data[0:i]), nil
	}
	if atEOF {
		return len(data), dropCR(data), nil
	}
	return 0, nil, nil
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
