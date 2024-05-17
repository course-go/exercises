package katas

import (
	"github.com/course-go/exercises/02-katas/internal/katas/numbers"
)

// Switch iterates over numbers and based on two predicates either
// sums, ignores them or exits the loop.
// However, it currently seems to be stuck in the loop.
// Fit it while avoiding the "if" statement.
func Switch() {
	var number int
	var sum int
	for {
		switch {
		case numbers.ShouldBreak(number):
			break
		case numbers.ShouldAdd(number):
			sum += number
		}

		number++
	}

	numbers.ProcessSum(sum)
}

// Fallthrough goes decides whether day is a weekday or weekend.
// However, it does not seem to work for some days.
// Fix it.
func Fallthrough(day string) string {
	var part string
	switch day {
	case "Monday":
	case "Tuesday":
	case "Wednesday":
	case "Thursday":
	case "Friday":
		part = "weekday"
	case "Saturday", "Sunday":
		part = "weekend"
	}

	return part
}
