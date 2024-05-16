package slices

import (
	"os"
)

func GetPaths() []string {
	return []string{
		"paths",
		"to",
		"my",
		"files",
	}
}

func ProcessFile(file *os.File) {}
