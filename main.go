package main

import (
	"fmt"

	"example.com/blueprint/modules"
)

func main() {
	fmt.Println("Hello, World!")

	myTicket := modules.Ticket{
		ID:    123,
		Event: "A multicultural event",
	}

	myTicket.PrintEvent()
}
