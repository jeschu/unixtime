package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

const format = "2006-01-02T15:04:05.000000000Z07:00"

func main() {
	for _, arg := range os.Args[1:] {
		if f, err := strconv.ParseFloat(arg, 64); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		} else {
			fmt.Fprintf(os.Stdout, "%s -> %s\n", toTime(f), arg)
		}
	}
}

func toTime(f float64) string {
	sec, dec := math.Modf(f)
	return time.Unix(int64(sec), int64(dec*(1e9))).Format(format)
}
