package lift

import (
	"time"
)

func newLift(id string, floors []int) *Lift {
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
	id           string
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
				continue
			}
		}
	}()
}

func (l *Lift) Stop() {
	if l.running != nil {
		l.running <- false
	}
	l.status = stationary
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

type status string

const (
	stationary status = "stationary"
	descending status = "descending"
	ascending  status = "ascending"
)
