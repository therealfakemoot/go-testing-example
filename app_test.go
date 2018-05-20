package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func ReadLines(f *os.File) []string {
	f.Seek(0, 0)
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, f)

	if err != nil {

	}

	fullFile := buf.String()
	lines := strings.Split(fullFile, "\n")

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
	// tempFile, teardown, err := setupTests("fm")
	tempFile, _, err := setupTests("fm")

	// defer teardown()

	if err != nil {
		t.Error("Unable to open temp file.")
	}

	fm := FileManipulator{f: tempFile}

	t.Run("Append", func(t *testing.T) {
		s := "xxxxx\n"
		fm.Append(s)

		lines := ReadLines(tempFile)

		if lines[len(lines)-1] != "xxxxx" {
			t.Logf("Lines:%v\n", lines)
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
