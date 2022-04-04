package main

import (
	"fmt"
)

func main() {

	var activeMenu string = "MAIN"
	// execute all the preloading activities
	preloadDefaultCategories() // categories
	preloadShoppingList()      // shopping list

	menuHandler(&activeMenu)

}

func menuHandler(menu *string) {

	var choice rune

	for {
		showMenu(*menu)
		choice = getInput()

		if *menu == "REPORTS" {
			// REPORTS menu
			if choice == rune('3') {
				*menu = "MAIN"
				showMenu(*menu)
			} else if choice >= rune('1') && choice <= rune('2') {
				break
			} else {
				fmt.Println("Your choice must be between 1 and 3.")
			}
		} else {
			// MAIN menu
			if choice == rune('2') {
				*menu = "REPORTS"
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

	redirectTo(*menu, choice)

}
