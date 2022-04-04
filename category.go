package main

import "fmt"

type intKeyValuePair struct {
	key   int
	value string
}

var categories []intKeyValuePair
var mapCategoryItem map[int]string

func preloadDefaultCategories() {
	// initialize the categories slice
	categories = append(categories, intKeyValuePair{key: 0, value: "Household"})
	categories = append(categories, intKeyValuePair{key: 1, value: "Food"})
	categories = append(categories, intKeyValuePair{key: 2, value: "Drinks"})
}

/*
	function to show the screen for adding a category
*/
func showAddCategoryScreen() {

	var name string //name of new categoryb

	var verifiedInput intKeyValuePair

	// clear the screen
	clearScreen()

	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("+ Add a New Category                   +")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")

	categoriesList := getLookupValues(categories)

	// infinit loop #1: input prompt -> category 'name'
	for {
		// fmt.Println(categories)
		fmt.Println("What is the category name to add ?")
		fmt.Scanln(&name)
		if hasPassed, message := inputNotEmptyCheck(&name, "-"); !hasPassed {
			message = "No Input Found !"
			fmt.Println(message)
		} else if isUnique, message := inputIsUniqueByNameCheck(&name, categoriesList); !isUnique {
			// is not unique
			fmt.Println(message)
			// fmt.Printf("isUnique -> %t, name -> %s, message -> %s", isUnique, name, message)
		} else {
			// passed.  exit from this infinite loop
			verifiedInput.key = len(categoriesList)
			verifiedInput.value = name
			categories = append(categories, verifiedInput)
			fmt.Printf("New Category : %s is added at index %d\n", name, len(categoriesList))
			fmt.Println("")
			break
		}
		fmt.Println("")

		returnToPrevious("MAIN")
	}

}

/*
	function to print out the categories (key, value)
	as a lookup card
*/
func showCategoriesLookup() {
	fmt.Printf("( ")
	for i, c := range categories {

		if i == (len(categories) - 1) {
			// last slice element
			fmt.Printf("%d -> %s )\n", c.key, c.value)
		} else {
			// not the last slice element
			fmt.Printf("%d -> %s   ", c.key, c.value)
		}

	}
}

/*
	function to print out all the Categories
*/
func printAllCategories() {
	for _, c := range categories {
		fmt.Println(c)
	}
}

/*
	function to get the category by the key
*/
func getCategoryByKey(k int) string {

	outcome := ""

	for _, c := range categories {
		if c.key == k {
			outcome = c.value
		}
	}
	return outcome
}
