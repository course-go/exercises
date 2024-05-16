package katas

import "github.com/course-go/exercises/03-katas/internal/katas/errors"

// Check demonstates error checking.
// However, the checks are not done properly.
// Fix the error checks.
func Check() error {
	p := errors.NewProcessor()
	err := p.Run()
	if err.Error() == "processor is not initialized" {
		p.Initialize()
	}

	err = p.Run()
	if err == errors.ErrPortAlreadyUsed {
		err = p.RunOnRandomPort()
	}

	if err != nil {
		return err
	}

	return nil
}
