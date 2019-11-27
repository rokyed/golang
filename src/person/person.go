package person

type Person struct {
	FName string
	LName string
}

func (p Person) SupBoi() string {
	return "I am " + p.FName + " " + p.LName
}

func (p *Person) ChangeLastName(name string) {
	p.LName = name
}

func (p *Person) ChangeFirstName(name string) {
	p.FName = name
}
