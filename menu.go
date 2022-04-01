package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
  function to list the menu items
*/
func showMenu(menu string) {
	switch menu {
	default:
		showMainMenu()
	case "REPORTS":
		showReportsMenu()
	}

}

func showMainMenu() {
	// clear the screen (i.e. stdout)
	clearScreen()
	fmt.Println("===============================")
	fmt.Println("= Shopping List Application   =")
	fmt.Println("===============================")
	fmt.Println("1. View entire shopping list")
	fmt.Println("2. Generate Shopping List Report")
	fmt.Println("3. Add Item")
	fmt.Println("4. Modify Item")
	fmt.Println("5. Delete Item")
	fmt.Println("6. Print Current Data")
	fmt.Println("7. Add New Category Name")
	fmt.Println("0. Exit")
	fmt.Println("")
	fmt.Printf("Select your choice : ")
}

/*
	function to list the Generate Report menu items
*/
func showReportsMenu() {
	clearScreen()
	fmt.Println("-------------------------------")
	fmt.Println("- Generate Report             -")
	fmt.Println("-------------------------------")
	fmt.Println("1. Total Cost of each category.")
	fmt.Println("2. List of item by category.")
	fmt.Println("3. Main Menu")
	fmt.Println("")
	fmt.Printf("Choose your report : ")
}

/*
  function to check if the input (from user) is a menu item listed
*/
func check(menu *string, choice rune) (bool, string) {
	var message string
	var outcome bool
	if *menu == "REPORTS" {
		// "REPORTS"
		if choice == rune('3') {
			outcome = true
			message = "Return to Main menu."
			*menu = "MAIN"
		} else if choice >= rune('1') && choice <= rune('2') {
			outcome = true
			message = "Ok"
		} else {
			outcome = false
			message = "Your choice must be between 1 and 3."
		}

	} else {
		// "MAIN" menu
		if choice == rune('2') {
			outcome = true
			message = "Go into Reports menu."
			*menu = "REPORTS"
		} else if choice == rune('0') {
			outcome = true
			message = "Have a good day.  Bye-bye."
		} else if choice == rune('1') ||
			(choice >= rune('3') && choice <= rune('7')) {
			outcome = true
			message = "MAIN"
		} else {
			outcome = false
			message = "Your choice must be between 0 and 7."
		}
		//
	}

	return outcome, message
}

/*
  function to get the ascii value of the char input by user (via stdin)
*/
func getInput() rune {

	reader := bufio.NewReader(os.Stdin)
	selected, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	return selected
}

/*
	function to redirect to the respective
	data handlers
*/
func redirectTo(menu string, selected rune) {
	switch selected {
	case rune('1'):
		// fmt.Println("You have selected 1")
		showShoppingList()
	case rune('3'):
		// add item
		// fmt.Println("You have selected 3")
		showAddItemScreen()
	case rune('4'):
		fmt.Println("You have selected 4")
	case rune('5'):
		fmt.Println("You have selected 5")
	case rune('6'):
		fmt.Println("You have selected 6")
	case rune('7'):
		fmt.Println("You have selected 7")
	default:
		fmt.Println("Exited Application.")
	}
}
