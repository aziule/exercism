package twelve

import (
	"bytes"
	"fmt"
	"strings"
)

var days = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var prefixes = map[int]string{
	1:  "a",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
}

var presents = map[int]string{
	1:  "Partridge in a Pear Tree",
	2:  "Turtle Doves",
	3:  "French Hens",
	4:  "Calling Birds",
	5:  "Gold Rings",
	6:  "Geese-a-Laying",
	7:  "Swans-a-Swimming",
	8:  "Maids-a-Milking",
	9:  "Ladies Dancing",
	10: "Lords-a-Leaping",
	11: "Pipers Piping",
	12: "Drummers Drumming",
}

func Song() string {
	buff := bytes.NewBuffer(nil)

	for i := 1; i <= len(days); i++ {
		buff.WriteString(Verse(i))
		buff.WriteString("\n")
	}

	return buff.String()
}

func Verse(num int) string {
	buff := bytes.NewBufferString("On the ")
	buff.WriteString(days[num])
	buff.WriteString(" day of Christmas my true love gave to me: ")

	p := []string{}
	var prefix string

	for i := num; i > 0; i-- {
		if i == 1 && num > 1 {
			prefix = fmt.Sprintf("and %s", prefixes[i])
		} else {
			prefix = prefixes[i]
		}

		p = append(p, fmt.Sprintf("%s %s", prefix, presents[i]))
	}

	buff.WriteString(strings.Join(p, ", "))
	buff.WriteString(".")
	return buff.String()
}
