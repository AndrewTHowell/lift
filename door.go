package lift

import "time"

func newDoor() *door {
	return &door{status: closed}
}

type door struct {
	status doorStatus
}

type doorStatus int

const (
	closed doorStatus = iota
	closing
	opening
	open
)

func (d *door) openDoor() {
	switch d.status {
	case closed, closing:
		// TODO: Not safe for concurrency
		d.status = opening
		time.Sleep(2 * time.Second)
		d.status = open
	}
}

func (d *door) closeDoor() {
	switch d.status {
	case open, opening:
		// TODO: Not safe for concurrency
		d.status = closing
		time.Sleep(2 * time.Second)
		d.status = closed
	}
}
