package custody

import (
	"errors"
	"time"

	joltactivity "github.com/thetimhosking/jolt/enterprise/activity"
)

type Valid interface {
	OK() error
}

type Custody struct {
	Activity    joltactivity.Activity
	ID          int
	Description string
	StartDate   time.Time
}

func (c *Custody) OK() error {
	if c.Description == "" {
		return errors.New("description required")
	}
	if c.StartDate.After(time.Now()) {
		return errors.New("Custody start date cannot be in the future")
	}
	return nil
}

func (c *Custody) CustodyCaseAdd(d string) error {
	s := c.Description
	println(s)
	return nil
}
