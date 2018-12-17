package secret

var secretCode = map[uint]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

func Handshake(code uint) []string {
	var x uint = 0
	reverse := false
	result := []string{}

	for x <= code {
		if x == 0 {
			x = 1
		} else {
			x = x << 1
		}

		if (code & x) != 0 {
			if x == 16 {
				reverse = true
				continue
			}
			result = append(result, secretCode[x])
		}
	}

	if reverse {
		i := 0
		j := len(result) - 1
		for i < j {
			result[i], result[j] = result[j], result[i]
			i++
			j--
		}
	}

	return result
}
