package main

import (
	"fmt"
	"time"
)

func main() {
	defer Tracktime(time.Now()) //<--THIS

	time.Sleep(500 * time.Millisecond)
}
func Tracktime(pre time.Time) time.Duration {
	elapsed := time.Since(pre)
	fmt.Println("elapsed:", elapsed)

	return elapsed
}

//elapsed: 501.102708ms
