package main

import (
	"fmt"
	"strconv"
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

	verifiedInput := ItemInfo{}

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
		name = getInputString()

		if hasPassed, message := inputNotEmptyCheck(&name, "Name"); !hasPassed {
			fmt.Println(message)
		} else {
			// passed.  exit from this infinite loop
			break
		}
	}
	fmt.Println("")

	// input prompt -> category (integer value)
	for {
		fmt.Println("What category does it belongs to ?")
		// fmt.Println("(0 -> Household, 1 -> Food, 2 -> Drinks)")
		showCategoriesLookup()
		fmt.Scanln(&category)
		if isValid, message := inputIsValidLookupValueCheck(&category, "Category", &categories); !isValid {
			fmt.Println(message)
			fmt.Println("")
		} else {
			// passed. exit from this infinite loop
			verifiedInput.category, _ = strconv.Atoi(category)
			break
		}
	}

	fmt.Println("")

	// infinite loop #2: input prompt -> quantity
	for {
		fmt.Println("How many units are there ?")
		fmt.Scanln(&quantity)
		if isValid, message := inputIsValidIntegerValue(&quantity, true, "Units"); !isValid {
			fmt.Println(message)
		} else {
			// passed. exit from this infinite loop
			verifiedInput.quantity, _ = strconv.Atoi(quantity)
			break
		}
	}

	fmt.Println("")

	// infinite loop #3 : input prompt -> unit cost
	for {
		fmt.Println("How much does it cost per unit ($)?")
		fmt.Scanln(&unitCost)

		if isValid, message := inputIsValidFloatValueCheck(&unitCost, true, "Unit Cost"); !isValid {
			fmt.Println(message)
		} else {
			// passed.  exit from this infinite loop
			v, _ := strconv.ParseFloat(unitCost, 64)
			verifiedInput.unitCost = v
			break
		}
	}

	// add item into the Shopping List
	// fmt.Println(verifiedInput)
	shoppingCart[name] = verifiedInput

	fmt.Println("")
	fmt.Println("New item added.")
	fmt.Println("---------------")
	fmt.Printf("Category      : %s \n", getCategoryByKey(shoppingCart[name].category))
	fmt.Printf("Item          : %s \n", name)
	fmt.Printf("Units         : %d \n", shoppingCart[name].quantity)
	fmt.Printf("Unit Cost ($) : %0.2f \n", shoppingCart[name].unitCost)
	fmt.Println("")

	returnToPrevious("MAIN")

}
