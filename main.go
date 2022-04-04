package main

import (
	"fmt"
)

func main() {

	// execute all the preloading activities
	preloadDefaultCategories() // categories
	preloadShoppingList()      // shopping list

	// mainMenuHandler()
	mainMenuHandlerV2()

}

func mainMenuHandler() {
	var choice rune
	var canExit bool = false
	var activeMenu string = "MAIN"

	for !canExit {
		// show the menu
		showMenu(activeMenu)

		fmt.Printf("activeMenu -> %s\n", activeMenu)
		// get the ascii value of the input from user (via stdin)
		choice = getInput()
		valid, message := check(&activeMenu, choice)
		if valid && message == "Ok" {
			// break from this infinite loop
			// and call the associated function
			canExit = valid

		} else if valid && message != "Ok" && activeMenu == "REPORTS" {
			// '3' was selected
			// the activeMenu has been set to "REPORTS" by check()
			// do nothing in here.
			// activeMenu = "MAIN"

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

func mainMenuHandlerV2() {
	var activeMenu string = "MAIN"
	var choice rune

	for {
		showMenu(activeMenu)
		choice = getInput()

		if activeMenu == "REPORTS" {
			// REPORTS menu
			if choice == rune('3') {
				activeMenu = "MAIN"
				showMenu(activeMenu)
			} else if choice >= rune('1') && choice <= rune('2') {
				break
			} else {
				fmt.Println("Your choice must be between 1 and 3.")
			}
		} else {
			// MAIN menu
			if choice == rune('2') {
				activeMenu = "REPORTS"
				// showMenu(activeMenu)

			} else if choice == rune('0') {
				fmt.Println("Have a good day.  Bye-bye.")
				break
			} else if choice == rune('1') ||
				(choice >= rune('3') && choice <= rune('7')) {
				break
			} else {
				fmt.Println("Your choice must be between 0 and 7.")
			}
		}
	}

	redirectTo(activeMenu, choice)

}
