package main

import "fmt"

type intKeyValuePair struct {
	key   int
	value string
}

var categories []intKeyValuePair

func preloadDefaultCategories() {
	// initialize the categories slice
	categories = append(categories, intKeyValuePair{key: 0, value: "Household"})
	categories = append(categories, intKeyValuePair{key: 1, value: "Food"})
	categories = append(categories, intKeyValuePair{key: 2, value: "Drinks"})
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
