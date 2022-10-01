package lift

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func NewController(minFloor, maxFloor, numberOfLifts int) *Controller {
	c := &Controller{
		minFloor:     minFloor,
		maxFloor:     maxFloor,
		FloorButtons: make(map[int]*Button, maxFloor-minFloor+1),
		Lifts:        make([]*Lift, numberOfLifts),
	}
	defer c.start()
	for floor := minFloor; floor <= maxFloor; floor++ {
		c.FloorButtons[floor] = &Button{floor}
	}
	for i := range c.Lifts {
		c.Lifts[i] = newLift(toLetter(i))
	}
	return c
}

type Controller struct {
	minFloor, maxFloor int
	FloorButtons       map[int]*Button
	Lifts              []*Lift
}

func (c *Controller) Stop() {
	for _, lift := range c.Lifts {
		lift.Stop()
	}
}

func (c *Controller) Press(button *Button) {
	button.Press(c)
}

func (c *Controller) callLiftToFloor(floor int) {
	closestLift := c.Lifts[0]
	closestLiftDistance := c.Lifts[0].distanceFrom(floor)
	for _, lift := range c.Lifts {
		distance := lift.distanceFrom(floor)
		if distance < closestLiftDistance {
			closestLift = lift
			closestLiftDistance = distance
		}
	}
	closestLift.sendToFloor(floor)
}

func (c *Controller) start() {
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				c.report()
				continue
			}
		}
	}()
}

func (c Controller) report() {
	var str string
	for floorNumber := c.maxFloor; floorNumber >= c.minFloor; floorNumber-- {
		str += fmt.Sprintf("%2d|", floorNumber)

		for _, l := range c.Lifts {
			lift := *l

			str += " "
			if lift.currentFloor == floorNumber {
				str += lift.id
			} else if lift.status == ascending && lift.currentFloor == floorNumber-1 {
				str += "▲"
			} else if lift.status == descending && lift.currentFloor == floorNumber+1 {
				str += "▼"
			} else {
				str += " "
			}
		}
		str += "\n"
	}
	wipeStdout()
	fmt.Print(str)
}

const abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func toLetter(i int) string {
	return abc[i : i+1]
}

func wipeStdout() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
