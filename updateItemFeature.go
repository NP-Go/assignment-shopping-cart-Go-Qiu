package main

import (
	"fmt"
	"strconv"
	"strings"
)

func showUpdateItemScreen() {

	var name string
	var category string
	var quantity string
	var unitCost string

	verifiedInput := ItemInfo{}

	clearScreen()

	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("+ Modify Item                          +")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")

	// infinit loop #1: input prompt -> name
	for {
		fmt.Println("Which item do you want to modify ?")
		// fmt.Scanln(&name)
		name = getInputString()

		if hasPassed, message := inputNotEmptyCheck(&name, "Name"); !hasPassed {
			fmt.Println(message)
		} else if existed := inputExistCheck(shoppingCart, name); !existed {
			fmt.Println("Item is not found !")
			fmt.Println("")
		} else {
			// passed.  exit from this infinite loop
			break
		}
	}
	fmt.Println("")
	fmt.Printf("Current item name is %s -  ", name)
	fmt.Printf("Category is %s -  ", getCategoryByKey(shoppingCart[name].category))
	fmt.Printf("Units is %d -  ", shoppingCart[name].quantity)
	fmt.Printf("Unit Cost is $ %0.2f\n", shoppingCart[name].unitCost)
	fmt.Println("")

	// input prompt -> category (integer value)
	for {
		fmt.Printf("Select new Category : (Press 'ENTER' to accept %d)\n", shoppingCart[name].category)
		showCategoriesLookup()
		fmt.Scanln(&category)

		if len(strings.TrimSpace(category)) == 0 {
			category = strconv.Itoa(shoppingCart[name].category)
			fmt.Println(category)
		}

		if isValid, message := inputIsValidLookupValueCheck(&category, "Category", &categories); !isValid {
			category = ""
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
		fmt.Printf("Enter new Units : (Press 'ENTER' to accept %d)\n", shoppingCart[name].quantity)
		fmt.Scanln(&quantity)

		// empty value or 'ENTER' is pressed
		if len(strings.TrimSpace(quantity)) == 0 {
			quantity = strconv.Itoa(shoppingCart[name].quantity)
			fmt.Println(quantity)
		}

		if isValid, message := inputIsValidIntegerValue(&quantity, true, "Units"); !isValid {
			quantity = ""
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
		fmt.Printf("Enter new Unit Cost ($) : (Press 'ENTER' to accept %0.2f)\n", shoppingCart[name].unitCost)
		fmt.Scanln(&unitCost)

		// empty value or 'ENTER' is pressed
		if len(strings.TrimSpace(unitCost)) == 0 {
			unitCost = strconv.Itoa(int(shoppingCart[name].unitCost))
			fmt.Println(unitCost)
		}

		if isValid, message := inputIsValidFloatValueCheck(&unitCost, true, "Unit Cost"); !isValid {
			unitCost = ""
			fmt.Println(message)
			fmt.Println("")
		} else {
			// passed.  exit from this infinite loop
			v, _ := strconv.ParseFloat(unitCost, 64)
			verifiedInput.unitCost = v
			break
		}
	}

	shoppingCart[name] = verifiedInput
	fmt.Println("")

	returnToPrevious("MAIN")

	//
}

/*
	function to check if a specific item
	exist in the shopping list
*/
func inputExistCheck(m map[string]ItemInfo, name string) bool {
	outcome := false
	for k, _ := range m {
		if k == strings.TrimSpace(name) {
			outcome = true
		}

	}
	return outcome
}
