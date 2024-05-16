package katas

import (
	"os"

	"github.com/course-go/exercises/03-katas/internal/katas/slices"
)

// Initialization initializes a new slice and extracts data.
// However, it does not behave completely correctly.
// Fix the mistake.
func Initialization() {
	movies := slices.GetMovies()
	directors := make([]slices.Director, len(movies))
	for _, movie := range movies {
		directors = append(directors, movie.Director)
	}

	slices.ProcessDirectors(directors)
}

// Resources open multiple resources and frees them afterward.
// However, it is not efficient and does not scale well with paths count.
// Fix the inefficiency.
func Resources() error {
	for _, path := range slices.GetPaths() {
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()
		slices.ProcessFile(file)
	}

	return nil
}
