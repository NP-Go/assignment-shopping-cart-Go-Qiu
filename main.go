package main

import (
	"fmt"
)

func main() {
	var choice rune
	var validChoice bool = false
	var activeMenu string = "MAIN"

	for !validChoice {
		// show the menu
		showMenu(activeMenu)

		// get the ascii value of the input from user (via stdin)
		choice = getInput()
		valid, message := check(&activeMenu, choice)
		if valid && message == "Ok" {
			// break from this infinite loop
			// and call the associated function
			validChoice = valid

		} else if valid && message != "Ok" {
			// break from this infinite loop
			// and end.
			validChoice = valid
			fmt.Println(message)
			fmt.Println("")
		} else {
			fmt.Println(message)
			fmt.Println("")
		}

	}

	// act on the valid choice

	redirectTo(activeMenu, choice)
	//
}
