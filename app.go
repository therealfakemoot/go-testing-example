package main

import (
	"bytes"
	"io"
	"os"
)

// FileManipulator implements Appender and Prepender
type FileManipulator struct {
	f *os.File
}

// Prepend adds argument to the beginning of the file.
func (fm *FileManipulator) Prepend(s string) {
	var b bytes.Buffer
	fm.f.Seek(0, 0)
	b.WriteString(s)
	io.Copy(&b, fm.f)
	fm.f.Truncate(0)
	fm.f.WriteAt(b.Bytes(), 0)
	fm.f.Seek(0, 0)
}

// Append adds argument to the end of the file.
func (fm *FileManipulator) Append(s string) {
	var b bytes.Buffer
	fm.f.Seek(0, 0)
	io.Copy(&b, fm.f)
	b.WriteString(s)
	fm.f.Truncate(0)
	fm.f.WriteAt(b.Bytes(), 0)
	fm.f.Seek(0, 0)
}
