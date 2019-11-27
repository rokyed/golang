package betterPerson

import (
	person "github.com/rokyed/golang/src/person"
)

type BetterPerson struct {
	person.Person
	IsBetter bool
}

func (p *BetterPerson) ItIsBetter() {
	p.IsBetter = true
}

func (p BetterPerson) SupBoi() string {
	return "I am the better of " + p.FName + " " + p.LName
}
