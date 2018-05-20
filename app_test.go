package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func ScanLines(r io.Reader) []string {
	s := bufio.NewScanner(r)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func setupTests(prefix string) (*os.File, func(), error) {
	createdFile, err := ioutil.TempFile(os.TempDir(), prefix)

	teardown := func() {
		os.Remove(os.TempDir() + createdFile.Name())
	}
	createdFile.WriteString("=====\n")
	createdFile.Sync()

	return createdFile, teardown, err
}

func TestFileManipulator(t *testing.T) {
	tempFile, teardown, err := setupTests("fm")

	defer teardown()

	if err != nil {
		t.Error("Unable to open temp file.")
	}

	fm := FileManipulator{f: tempFile}

	t.Run("Append", func(t *testing.T) {
		s := "xxxxx\n"
		fm.Append(s)

		lines := ScanLines(tempFile)

		if lines[len(lines)-1] != "xxxxx" {
			t.Logf("Lines:%#v\n", lines)
			t.Fail()
		}

	})

	t.Run("Prepend", func(t *testing.T) {
		s := "yyyyy\n"
		fm.Prepend(s)

		lines := ScanLines(tempFile)

		if lines[0] != "yyyyy" {
			t.Logf("Lines:%#v\n", lines)
			t.Fail()
		}

	})
}
