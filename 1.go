package main

import (
	"fmt"

	betterPerson "github.com/rokyed/golang/src/betterPerson"
	person "github.com/rokyed/golang/src/person"
)

func main() {
	p := person.Person{FName: "Ali", LName: "G"}
	p1 := betterPerson.BetterPerson{Person: p, IsBetter: true}

	fmt.Println(p1.SupBoi())
	p1.ChangeLastName("G-abriel")
	fmt.Println(p1.SupBoi())
}
