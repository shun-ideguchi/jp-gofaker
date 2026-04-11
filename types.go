package jpfaker

type PersonName struct {
	LastName      string
	FirstName     string
	LastNameKana  string
	FirstNameKana string
}

func (n PersonName) FullName() string {
	return n.LastName + " " + n.FirstName
}

func (n PersonName) FullNameKana() string {
	return n.LastNameKana + " " + n.FirstNameKana
}

type AddressValue struct {
	PostalCode string
	Prefecture string
	City       string
	Street     string
	Building   string
}

func (a AddressValue) Full() string {
	return a.Prefecture + a.City + a.Street + a.Building
}
