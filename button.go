package lift

type Button struct {
	floor int
}

func (b *Button) Press(c *Controller) {
	c.callLiftToFloor(b.floor)
}
