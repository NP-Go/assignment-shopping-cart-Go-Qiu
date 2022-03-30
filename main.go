package main

import (
	"bufio"
	"fmt"
	"os"
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
	//
}

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
