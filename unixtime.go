package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

const format = "2006-01-02T15:04:05.000000000Z07:00"

func main() {
	reg, err := regexp.Compile(`[^0-9\.]+`)
	if err != nil {
		log.Fatal(err)
	}
	for _, arg := range os.Args[1:] {
		v := reg.ReplaceAllString(arg, "")
		if f, err := strconv.ParseFloat(v, 64); err != nil {
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
