package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	unix := t.Unix()
	unixNano := t.UnixNano()
	fmt.Printf("t: %v\n", t)
	fmt.Println(year, month, day, hour, minute, second, unix, unixNano)
	fmt.Printf("time.Unix(unix): %v\n", time.Unix(0, unixNano))
	fmt.Printf("t.Format(\"2006-01-02 15:04:05\"): %v\n", t.Format("2006-01-02 15:04:05"))
}
