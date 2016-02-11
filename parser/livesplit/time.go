package livesplit

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// parseTime takes a LiveSplit duration and returns the raw number of milliseconds it represents. A LiveSplit duration
// is a string that looks like
//
//   (\d\d:)?(\d\d:)?(\d\d:)?((\d\d\)(\.\d+)?)?
//
// where \1 is the number of days, \2 is the number of hours, \3 is the number of minutes, and \4 is the number of
// seconds as a possible floating point number. Example LiveSplit times include:
//
//   "00:00:02.415"   // 2.415 seconds
//   "02:10:00"       // 2 hours and 10 minutes
//   "12.4"           // 12.4 seconds
//   "01:01:01:01.01" // 1 day, 1 hour, 1 minute, and 1.01 seconds
//   ""               // 0 seconds
//
// If the given time is more accurate than milliseconds, the additional accuracy is ignored. No rounding takes place.
func parseTime(t string) (time.Duration, error) {
	if t == "" {
		return 0, nil
	}

	s := strings.Split(t, ":")
	l := len(s)

	if l < 1 {
		return 0, errors.New("can't parse LiveSplit file: couldn't split time on ':'")
	}

	var d time.Duration

	// Milliseconds and seconds
	if l >= 1 {
		seconds, err := strconv.ParseFloat(s[l-1], 64)
		milliseconds := seconds * 1000
		if err != nil {
			return 0, fmt.Errorf("couldn't convert seconds to float: %s", err)
		}
		d += time.Duration(milliseconds) * time.Millisecond
	}

	// Minutes
	if l >= 2 {
		minutes, err := strconv.Atoi(s[l-2])
		if err != nil {
			return 0, fmt.Errorf("couldn't convert minutes to int: %s", err)
		}
		d += time.Duration(minutes) * time.Minute
	}

	// Hours
	if l >= 3 {
		hours, err := strconv.Atoi(s[l-3])
		if err != nil {
			return 0, fmt.Errorf("couldn't convert hours to int: %s", err)
		}
		d += time.Duration(hours) * time.Hour
	}

	// Days
	if l >= 4 {
		days, err := strconv.Atoi(s[l-4])
		if err != nil {
			return 0, fmt.Errorf("couldn't convert days to int: %s", err)
		}
		d += time.Duration(days*24) * time.Hour
	}

	return d, nil
}
