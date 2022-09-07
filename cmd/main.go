package main

import (
	"time"

	"lift"
)

func main() {
	l := lift.NewLift(1)
	defer l.Stop()

	time.Sleep(5 * time.Second)
}
