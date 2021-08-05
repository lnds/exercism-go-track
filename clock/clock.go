package clock

import "fmt"

type Clock struct {
	minutes int
}

func New(hours int, minutes int) Clock {
	clock := Clock{minutes: ((hours*60+minutes)%1440 + 1440) % 1440}
	return clock
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.minutes/60, clock.minutes%60)
}

func (clock Clock) Add(minutes int) Clock {
	return New(clock.minutes/60, clock.minutes%60+minutes)
}

func (clock Clock) Subtract(minutes int) Clock {
	return clock.Add(-minutes)
}
