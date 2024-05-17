package katas

import (
	"os"

	"github.com/course-go/exercises/02-katas/internal/katas/files"
)

// Defer closes an openned file.
// Does the code behave as expected?
func Defer() {
	// The file struct represent a file handle.
	// It is further explored in later lectures.
	var file os.File
	defer file.Close()
	file = files.OpenedFile()
}
