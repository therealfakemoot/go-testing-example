package main

import (
	"io"
)

// Appender represents a value that can write bytes to the end of a buffer.
type Appender interface {
	Append([]byte) (int, error)
}

// Prepender represents a value that can write bytes to the start of a buffer.
type Prepender interface {
	Prepend([]byte) (int, error)
}

// FileManipulator implements Appender and Prepender
type FileManipulator struct {
	f io.ReadWriteSeeker
}

// Prepend adds argument to the beginning of the file.
func (fm *FileManipulator) Prepend(s string) {}

// Append adds argument to the end of the file.
func (fm *FileManipulator) Append(s string) {
}
