package main

import (
	"fmt"
)

func main() {
	var myName string = "Sergio Terroso"
	myAge := 46
	myPi := 3.1416

	fmt.Printf("Hello, my name is %s, I'm %d years old and PI is %.2f\n", myName, myAge, myPi)

	var myNonInitializedString string
	var myNonInitializedInteger int
	var myNonInitializedBoolean bool
	var myNonInitializedFloat float32

	fmt.Printf(
		"My non-initialized variables:\nString: '%s'\nInteger: %d\nBoolean: %t\nFloat: %f\n",
		myNonInitializedString,
		myNonInitializedInteger,
		myNonInitializedBoolean,
		myNonInitializedFloat)
}
