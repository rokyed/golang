package person

import (
	"fmt"
	"strconv"
)

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

func (p Person) SayHello(ammount int) {
	for i := 0; i < ammount; i++ {
		fmt.Println("Hello, my name is " + p.FName + " " + p.LName)
	}
}

func (p Person) SayHelloWhile(ammount int) {
	i := 0
	for {
		if i >= ammount {
			break
		}

		fmt.Println("Hello, my name is " + p.FName + " " + p.LName)

		i++
	}
}

func (p Person) PlayFizzBuzz(ammount int, skip int) {
	i := 1
	for {
		if i == skip {
			fmt.Println(" ")
			i++
			continue
		}

		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}

		if i == ammount {
			break
		}

		i++
	}
	txt := "Did you realized we skiddaddled "
	txt += strconv.Itoa(skip)
	txt += "? no ? look back :)"
	fmt.Println(txt)
}
