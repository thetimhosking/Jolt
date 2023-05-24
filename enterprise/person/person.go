package person

import "time"

type Person struct {
	Id int64
}

type PersonName struct {
	Id            int64
	PersonID      int64
	FirstName     string
	MiddleName    string
	LastName      string
	NameSuffix    string
	NameTitle     string
	ValidFromDate time.Date
	ValidToDate   time.Date
}

func (p *Person) Validate() (bool, error) {

}
