package main

import (
	"fmt"
	"strings"
)

/*
	function to show the screen for
	getting the data of a new item
*/
func showAddItemScreen() {

	// variables to capture the input from user (via stdin)
	name := ""
	category := ""
	quantity := ""
	unitCost := ""

	// continueToLoop := true

	// clear the screen
	clearScreen()

	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("+ Add a New Item                       +")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")

	// infinit loop #1: input prompt -> name
	for {
		fmt.Println("What is the name of the NEW item ?")
		fmt.Scanln(&name)
		if hasPassed, message := inputNotEmptyCheck(&name, "Name"); !hasPassed {
			fmt.Println(message)
		} else {
			// passed.  exit from this infinite loop
			break
		}
	}

	// input prompt -> category (integer value)
	fmt.Println("What category does it belongs to ?")
	fmt.Println("(0 -> Household, 1 -> Food, 2 -> Drinks)")
	fmt.Scanln(&category)

	// infinite loop #2: input prompt -> quantity
	for {
		fmt.Println("How many units are there ?")
		fmt.Scanln(&quantity)
		if isValid, message := inputIsValidIntegerValue(&quantity, true, "Units"); !isValid {
			fmt.Println(message)
		} else {
			// passed. exit from this infinite loop
			break
		}
	}

	// infinite loop #3 : input prompt -> unit cost
	for {
		fmt.Println("How much does it cost per unit ($)?")
		fmt.Scanln(&unitCost)

		if isValid, message := inputIsValidFloatValue(&unitCost, true, "Unit Cost"); !isValid {
			fmt.Println(message)
		} else {
			// passed.  exit from this infinite loop
			break
		}
	}
}

/*
	function to execute checks on the specified entries
	to ensure that input (from user via stdin) matches
	the defined rules.
*/
func inputNotEmptyCheck(input *string, inputName string) (bool, string) {
	var tag string
	if len(strings.TrimSpace(inputName)) != 0 {
		tag = strings.TrimSpace(inputName)
	} else {
		tag = "Input value"
	}
	if len(strings.TrimSpace(*input)) == 0 {
		// empty string
		return false, tag + " cannot be empty."
	} else {
		// non-empty string
		// exit from infinite loop

		return true, "Ok"
	}
}

/*
	SUPPORTING FUNCTIONS SECTION
*/
