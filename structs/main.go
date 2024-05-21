package main

import "fmt"

type Person struct {
	FirstName  string
	MiddleName string
	LastName   string
	IsWalking  bool
}

func PersonFactory(firstName string, middleName string, lastName string, isWalking bool) Person {
	return Person{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		IsWalking:  isWalking,
	}
}

func (p *Person) SetFirstName(firstName string) {
	p.FirstName = firstName
}

func main() {
	myPerson := PersonFactory("Sergio", "", "Terroso Cabrera", false)

	var walkingState string

	myPerson.IsWalking = true

	myPerson.SetFirstName("Silvia")

	if myPerson.IsWalking {
		walkingState = "is"
	} else {
		walkingState = "isn't"
	}

	fmt.Printf("My person's First Name is '%s'\n", myPerson.FirstName)
	fmt.Printf("My person's Middle Name is '%s'\n", myPerson.MiddleName)
	fmt.Printf("My person's Last Name is '%s'\n", myPerson.LastName)
	fmt.Printf("My person %s walking\n", walkingState)
}
