package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
  function to list the menu items
*/
func showMenu() {

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
  function to check if the input (from user) is a menu item listed
*/
func check(choice rune) (bool, string) {
	var message string
	var outcome bool

	if choice > rune('0') && choice <= rune('7') {
		outcome = true
		message = "Ok"
	} else if choice == rune('0') {
		outcome = true
		message = "Have a good day.  Bye-bye."
	} else {
		outcome = false
		message = "Your choice must be between 0 and 7."
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
func redirect(selected rune) {
	switch selected {
	case rune('1'):
		fmt.Println("You have selected 1")

	case rune('2'):
		fmt.Println("You have selected 2")
	case rune('3'):
		fmt.Println("You have selected 3")
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
