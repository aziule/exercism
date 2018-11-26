package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
    freq := FreqMap{}
    freqMapQueue := make(chan FreqMap)
	defer close(freqMapQueue)

	for _, text := range texts {
	    go func(text string) {
	        freqMapQueue <-Frequency(text)
        }(text)
    }

	for range texts {
        for k, v := range <-freqMapQueue {
            freq[k] += v
        }
    }

    return freq
}
