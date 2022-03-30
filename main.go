package main

import (
	"fmt"
)

func main() {
	var choice rune
	var validChoice bool = false

	for !validChoice {
		// show the menu
		showMenu()

		// get the ascii value of the input from user (via stdin)
		choice = getInput()

		ok, message := check(choice)
		if ok && message == "Ok" {
			// break from this infinite loop
			// and call the associated function
			validChoice = ok

		} else if ok && message != "Ok" {
			// break from this infinite loop
			// and end.
			validChoice = ok
			fmt.Println(message)
			fmt.Println("")
		} else {
			fmt.Println(message)
			fmt.Println("")
		}

	}

	// act on the valid choice
	redirect(choice)
	//
}
