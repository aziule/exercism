package gigasecond

import "time"

const Gigasecond time.Duration = 1e9 * time.Second

// AddGigasecond adds one Gigasecond to a given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
