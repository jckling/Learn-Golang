package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	// current time
	t := time.Now()
	fmt.Println(t)			// 2021-04-19 21:57:14.2135836 +0800 CST m=+0.004295601
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())		// 19.04.2021

	t = time.Now().UTC()
	fmt.Println(t)          // 2021-04-19 13:57:14.2702794 +0000 UTC
	fmt.Println(time.Now()) // 2021-04-19 21:57:14.2702794 +0800 CST m=+0.060991401

	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 	// must be in nanosec
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now) 		// 2021-04-26 13:57:14.2702794 +0000 UTC

	// formatting times:
	fmt.Println(t.Format(time.RFC822)) // 19 Apr 21 13:57 UTC
	fmt.Println(t.Format(time.ANSIC))  // Mon Apr 19 13:57:14 2021

	// The time must be 2006-01-02 15:04:05
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 19 Apr 2021 13:57
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	// 2021-04-19 13:57:14.2702794 +0000 UTC => 20210419
}
