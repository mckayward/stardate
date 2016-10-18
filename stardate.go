package stardate

import (
	"fmt"
	"time"
)

const (
	SecondsPerStardate = 34367.0564
	UnixOfStardateZero = 10997841600 // July 5, 2315 at 12PM PST
)

type Stardate float64

func (s Stardate) String() string {
	return fmt.Sprintf("%06.1f", s)
}

func TimeToStardate(t time.Time) Stardate {
	unix_time := t.Unix()

	unix_diff := float64(unix_time - UnixOfStardateZero)

	s := Stardate(unix_diff / SecondsPerStardate)
	return s
}

func StardateToTime(s Stardate) time.Time {
	unix_of_stardate := s*SecondsPerStardate + UnixOfStardateZero

	return time.Unix(int64(unix_of_stardate), 0)
}
