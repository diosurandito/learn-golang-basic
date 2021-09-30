package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	layout := "2006-01-02"
	parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	parseWithLayout, _ := time.Parse(layout, "2020-01-02")
	fmt.Println(parse)
	fmt.Println(parseWithLayout)

}
