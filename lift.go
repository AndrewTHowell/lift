package lift

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func NewLift(id int, minFloor, maxFloor int) *Lift {
	floors := make([]int, maxFloor-minFloor+1)
	for i := range floors {
		floors[i] = minFloor + i
	}

	lift := &Lift{
		id:           id,
		currentFloor: 0,
		floors:       floors,
		status:       stationary,
		door:         newDoor(),
	}
	lift.Start()
	return lift
}

type Lift struct {
	id           int
	currentFloor int
	targetFloor  int // TODO: more than one target floor
	floors       []int
	status
	*door
	running chan bool
}

func (l *Lift) Start() {
	l.Stop()

	l.running = make(chan bool)
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-l.running:
				return
			case <-ticker.C:
				l.move()
				l.report()
				continue
			}
		}
	}()
	return
}

func (l *Lift) Stop() {
	if l.running != nil {
		l.running <- false
	}
}

func (l *Lift) SetTargetFloor(floorNumber int) {
	if l.floors[0] <= floorNumber && floorNumber <= l.floors[len(l.floors)-1] {
		l.targetFloor = floorNumber
	}
}

func (l *Lift) move() {
	vector := l.targetFloor - l.currentFloor
	if vector == 0 {
		l.status = stationary
		return
	}

	if l.status != stationary {
		if vector > 0 {
			l.currentFloor++
		} else if vector < 0 {
			l.currentFloor--
		}
	}

	if vector > 0 {
		l.status = ascending
	} else if vector < 0 {
		l.status = descending
	}

	nextVector := l.targetFloor - l.currentFloor
	if nextVector == 0 {
		l.status = stationary
	}
}

func (l *Lift) report() {
	liftColumn := make([]string, len(l.floors))
	for i, floor := range l.floors {
		if l.currentFloor == floor {
			liftColumn[i] = fmt.Sprint(l.id)
			if l.status == ascending {
				liftColumn[i+1] = "▲"
			} else if l.status == descending {
				liftColumn[i-1] = "▼"
			}
		}
	}

	var str string
	for i := len(liftColumn) - 1; i > -1; i-- {
		str += fmt.Sprintf("%2d| %s\n", l.floors[i], liftColumn[i])
	}
	wipeStdout()
	fmt.Print(str)
}

type status string

const (
	stationary status = "stationary"
	descending status = "descending"
	ascending  status = "ascending"
)

func wipeStdout() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
