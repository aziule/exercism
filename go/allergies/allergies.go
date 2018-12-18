package allergies

import (
	"sort"
)

var substances = map[int]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

// Allergies returns the list of allergies given a score
func Allergies(score uint) []string {
	var allergies []string
	var keys []int

	// Iterating over maps is randomised, but the expected result is sorted,
	// so we need to sort the keys beforehand
	for k := range substances {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		if score&uint(k) == uint(k) {
			allergies = append(allergies, substances[k])
		}
	}

	return allergies
}

// AllergicTo tells if a person is allergic to a substance given a score
func AllergicTo(score uint, substance string) bool {
	for k, v := range substances {
		if v == substance {
			return score&uint(k) == uint(k)
		}
	}

	return false
}
