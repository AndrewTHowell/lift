package main

import (
	"time"

	"lift"
)

func main() {
	l := lift.NewLift(1, -1, 5)
	defer l.Stop()

	l.SetTargetFloor(5)

	time.Sleep(7 * time.Second)

	l.SetTargetFloor(-1)

	time.Sleep(7 * time.Second)
}
