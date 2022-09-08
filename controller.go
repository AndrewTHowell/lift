package lift

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func NewController(minFloor, maxFloor int) *Controller {
	floors := make([]int, maxFloor-minFloor+1)
	for i := range floors {
		floors[i] = minFloor + i
	}

	controller := &Controller{
		floors: floors,
	}
	controller.start()
	return controller
}

type Controller struct {
	floors []int
	lifts  []*Lift
}

func (c *Controller) AddLift() *Lift {
	lift := newLift(toLetter(len(c.lifts)), c.floors)
	c.lifts = append(c.lifts, lift)
	return lift
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
	for i := len(c.floors) - 1; i > -1; i-- {
		floor := c.floors[i]
		str += fmt.Sprintf("%2d|", floor)

		for _, l := range c.lifts {
			lift := *l

			str += " "
			if lift.currentFloor == floor {
				str += lift.id
			} else if lift.status == ascending && lift.currentFloor == floor-1 {
				str += "▲"
			} else if lift.status == descending && lift.currentFloor == floor+1 {
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
	return abc[i-1 : i]
}

func wipeStdout() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
