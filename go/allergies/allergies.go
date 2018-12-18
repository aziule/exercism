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

func Allergies(score uint) []string {
	var allergies []string
	var keys []int

	for k := range substances {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		if score&uint(k) != 0 {
			allergies = append(allergies, substances[k])
		}
	}

	return allergies
}

func AllergicTo(score uint, substance string) bool {
	for k, v := range substances {
		if v == substance {
			return score&uint(k) != 0
		}
	}

	return false
}
