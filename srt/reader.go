package srt

import (
	"io"
	"os"
	"path/filepath"
)

func Open(path string) ([]Subtitle, error) {
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Read(f)
}

func Read(r io.Reader) ([]Subtitle, error) {
	s := newScanner(r)
	subtitles := make([]Subtitle, 0)
	for s.Scan() {
		subtitles = append(subtitles, s.Subtitle())
	}
	err := s.Err()
	if err != nil {
		return nil, err
	}

	return subtitles, nil
}
