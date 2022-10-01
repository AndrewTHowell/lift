package main

import (
	"time"

	"lift"
)

func main() {
	controller := lift.NewController(-1, 5, 2)
	defer controller.Stop()

	controller.Press(controller.FloorButtons[2])

	time.Sleep(5 * time.Second)

	controller.Press(controller.FloorButtons[-1])

	time.Sleep(3 * time.Second)

	controller.Press(controller.FloorButtons[5])

	time.Sleep(5 * time.Second)

	controller.Press(controller.FloorButtons[3])

	time.Sleep(5 * time.Second)
}
