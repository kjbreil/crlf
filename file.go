package crlf

import (
	"io"
	"os"
)

type writeCloser struct {
	io.Writer
	io.Closer
}

// Create opens a text file for writing, with platform-appropriate line ending
// conversion.
func Create(name string) (io.WriteCloser, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	return writeCloser{
		Writer: NewWriter(f),
		Closer: f,
	}, nil
}
