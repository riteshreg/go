package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	now := time.Now()

	p(now)

	then := time.Date(2004, 01, 29, 19, 15, 30, 651387237, time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Date())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(now.After(then))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
}
