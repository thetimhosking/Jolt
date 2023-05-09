package custody

import (
	"errors"
	"time"
)

type Valid interface {
	OK() error
}

type Custody struct {
	ID          int
	Description string
	StartDate   time.Time
}

func (c *Custody) OK() error {
	if c.Description == "" {
		return errors.New("Description required")
	}
	if c.StartDate.After(time.Now()) {
		return errors.New("Custody start date cannot be in the future")
	}
	return nil
}

func (c *CustodyCase) CustodyCaseAdd(d string) error {
	s := c.Description
	println(s)
	return nil
}