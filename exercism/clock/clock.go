package clock

import (
	"math"
	"strconv"
)

//Clock type definition
type Clock struct {
	h int
	m int
}

//New creates a new Clock
func New(h, m int) Clock {
	tempClock := Clock{h, m}
	tempMins := tempClock.ClockToMins()
	return MinsToClock(tempMins)
}

//Add : adds m minutes to a Clock
func (c Clock) Add(m int) Clock {
	tempMins := c.ClockToMins() + m
	return MinsToClock(tempMins)
}

//Subtract: subtracts m minutes from a clock
func (c Clock) Subtract(m int) Clock {
	tempMins := c.ClockToMins() - m
	return MinsToClock(tempMins)
}

//String: converts the clock type to a human readable string i.e. HH:MM
func (c Clock) String() string {
	var result string
	if c.h < 10 {
		result += "0" + strconv.Itoa(c.h) + ":"
	} else {
		result += strconv.Itoa(c.h) + ":"
	}

	if c.m < 10 {
		result += "0" + strconv.Itoa(c.m)
	} else {
		result += strconv.Itoa(c.m)
	}

	return result
}

// helper function to convert clock to minutes
func (c Clock) ClockToMins() int {
	return (c.h * 60) + c.m
}

//helper function to convert minutes to clock
func MinsToClock(m int) Clock {
	//discard days i.e. reduce to + or - 24 hours
	//m = m % 1440 !! golang modulo operator works different than other languages!!
	m = ((m%1440 + 1440) % 1440)

	h := int(math.Abs(float64(m / 60)))
	h = int(math.Abs(float64(h % 24)))

	m = int(math.Abs(float64(m % 60)))
	return Clock{h, m}
}
