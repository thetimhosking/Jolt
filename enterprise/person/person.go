package person

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
	ValidFromDate string
	ValidToDate   string
}

func (p *Person) Validate() (bool, error) {
	return true, nil
}
