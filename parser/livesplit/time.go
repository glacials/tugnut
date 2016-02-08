package livesplit

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

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

	if l >= 1 {
		seconds, err := strconv.ParseFloat(s[l-1], 64)
		milliseconds := seconds * 1000
		if err != nil {
			return 0, fmt.Errorf("couldn't convert seconds to float: %s", err)
		}
		d += time.Duration(milliseconds) * time.Millisecond
	}

	if l >= 2 {
		minutes, err := strconv.Atoi(s[l-2])
		if err != nil {
			return 0, fmt.Errorf("couldn't convert minutes to int: %s", err)
		}
		d += time.Duration(minutes) * time.Minute
	}

	if l >= 3 {
		hours, err := strconv.Atoi(s[l-3])
		if err != nil {
			return 0, fmt.Errorf("couldn't convert hours to int: %s", err)
		}
		d += time.Duration(hours) * time.Hour
	}

	if l >= 4 {
		days, err := strconv.Atoi(s[l-4])
		if err != nil {
			return 0, fmt.Errorf("couldn't convert days to int: %s", err)
		}
		d += time.Duration(days*24) * time.Hour
	}

	return d, nil
}
