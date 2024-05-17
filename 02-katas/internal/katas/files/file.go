package files

import "os"

func OpenedFile() os.File {
	file, _ := os.CreateTemp("", "dummy")
	return *file
}
