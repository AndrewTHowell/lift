package main

import (
	"time"

	"lift"
)

func main() {
	controller := lift.NewController(-1, 5)

	lift1 := controller.AddLift()
	defer lift1.Stop()
	lift2 := controller.AddLift()
	defer lift2.Stop()

	lift1.SetTargetFloor(5)

	time.Sleep(1 * time.Second)

	lift2.SetTargetFloor(-1)

	time.Sleep(7 * time.Second)

	lift1.SetTargetFloor(-1)
	lift2.SetTargetFloor(3)

	time.Sleep(8 * time.Second)
}
