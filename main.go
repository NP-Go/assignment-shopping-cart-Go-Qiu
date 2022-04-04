package main

import (
	"fmt"
)

func main() {
	var choice rune
	var canExit bool = false
	var activeMenu string = "MAIN"

	// execute all the preloading activities
	preloadDefaultCategories() // categories
	preloadShoppingList()      // shopping list

	for !canExit {
		// show the menu
		showMenu(activeMenu)

		// get the ascii value of the input from user (via stdin)
		choice = getInput()
		valid, message := check(&activeMenu, choice)
		if valid && message == "Ok" {
			// break from this infinite loop
			// and call the associated function
			canExit = valid

		} else if valid && message != "Ok" && activeMenu == "REPORTS" {
			// '2' was selected
			// show the Reports menu items
			showMenu(activeMenu)

		} else if valid && message != "Ok" {
			// '0' was selected
			// break from this infinite loop
			// and end.
			canExit = valid

		} else {
			fmt.Println(message)
			fmt.Println("")
		}

	}

	// act on the valid choice
	// fmt.Println(activeMenu)
	redirectTo(activeMenu, choice)
	//
}
