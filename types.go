package jpfaker

// PersonName represents a Japanese personal name with kana.
type PersonName struct {
	LastName      string
	FirstName     string
	LastNameKana  string
	FirstNameKana string
}

// FullName returns the surname and given name joined by a space.
func (n PersonName) FullName() string {
	return n.LastName + " " + n.FirstName
}

// FullNameKana returns the kana surname and given name joined by a space.
func (n PersonName) FullNameKana() string {
	return n.LastNameKana + " " + n.FirstNameKana
}

// AddressValue represents a generated Japanese postal address.
type AddressValue struct {
	PostalCode string
	Prefecture string
	City       string
	Street     string
	Building   string
}

// Full returns the full address as a single string.
func (a AddressValue) Full() string {
	return a.Prefecture + a.City + a.Street + a.Building
}
