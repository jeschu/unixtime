package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

const formatLocal = "2006-01-02T15:04:05.000000000Z07:00"
const formatUtc = "2006-01-02T15:04:05.000000000Z"

func main() {
	reg := regexp.MustCompile(`[^0-9\.]+`)

	for _, arg := range os.Args[1:] {
		v := reg.ReplaceAllString(arg, "")
		if f, err := strconv.ParseFloat(v, 64); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		} else {
			fmt.Fprintf(os.Stdout, "%s [%s] -> %s\n", toTime(f).Format(formatLocal), toTime(f).UTC().Format(formatUtc), arg)
		}
	}
}

func toTime(f float64) time.Time {
	sec, dec := math.Modf(f)
	return time.Unix(int64(sec), int64(dec*(1e9)))
}
