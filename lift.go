package lift

import (
	"fmt"
	"time"
)

func NewLift(id int) *Lift {
	lift := &Lift{
		id: id,
		currentFloor: 0,
		status:       stationary,
		door:         newDoor(),
	}
	lift.Start()
	return lift
}

type Lift struct {
	id int
	currentFloor int
	targetFloor  int // TODO: more than one target floor
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
				// TODO: stay, move up or down
				l.Report()
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
	l.targetFloor = floorNumber
}

func (l *Lift) Report() {
	fmt.Println(fmt.Sprintf("Lift %d, currently %s on floor %d", l.id, l.status, l.currentFloor))
}

type status string

const (
	stationary status = "stationary"
	// descending status = "descending"
	// ascending status = "ascending"
)
