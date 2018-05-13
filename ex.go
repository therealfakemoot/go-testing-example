package main

import "os"

// Clean removes a given file/directory. If there is an error, it is of type *PathError
func Clean(path string) error {
	return os.Remove(path)
}
