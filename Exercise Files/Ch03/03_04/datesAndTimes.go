package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)

	now := time.Now()
	fmt.Printf("The time now is %s\n", now)

	fmt.Println("The month is", t.Month())
	fmt.Println("The day is", t.Day())
	fmt.Println("The weekday is", t.Weekday())
}
