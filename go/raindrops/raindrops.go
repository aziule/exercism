package raindrops

import (
    "bytes"
    "sort"
    "strconv"
)

var factorsMap = map[int]string {
    3: "Pling",
    5: "Plang",
    7: "Plong",
}

// Convert converts a number into a string given the rules defined in factorsMap.
// Returns the number as a string if no factor matches.
func Convert(number int) string {
    hasFactor := false
    buff := bytes.Buffer{}
    factors := getFactorsOrdered()

    for _, factor := range factors {
        if number % factor == 0 {
            hasFactor = true
            buff.WriteString(factorsMap[factor])
        }
    }

    if !hasFactor {
        buff.WriteString(strconv.Itoa(number))
    }

    return buff.String()
}

func getFactorsOrdered() []int {
    var factors []int

    for factor := range factorsMap {
        factors = append(factors, factor)
    }

    sort.Ints(factors)

    return factors
}
