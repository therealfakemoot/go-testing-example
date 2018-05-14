package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func setupTests(prefix string) (string, func(), error) {
	createdFile, err := ioutil.TempFile(os.TempDir(), prefix)

	teardown := func() {
		os.Remove(os.TempDir() + createdFile.Name())
	}

	return createdFile.Name(), teardown, err
}

func TestFileCleanup(t *testing.T) {
	filename, teardown, err := setupTests("testA")

	if err != nil {
		t.Log("Failed to create temporary file.")
		t.Skip()
	}

	defer teardown()

	t.Run("testA", func(t *testing.T) {

		err = Clean(filename)

		if err != nil {
			t.Error("Failed to delete file.")
		}
	})
}
