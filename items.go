package main

import "fmt"

//  enumeration for Categories
const (
	Household = iota
	Food
	Drinks
)

// struct for Item data
type Item struct {
	category int
	name     string
	quantity int
	unitCost float64
}

// declare a slice to cache the items data
var shoppingList = []Item{}

/*
	function to preload the data in the shopping list
*/
func preloadShoppingList() {

	shoppingList = append(shoppingList, Item{category: 0, name: "Fork", quantity: 4, unitCost: 3.00})
	shoppingList = append(shoppingList, Item{category: 0, name: "Plates", quantity: 4, unitCost: 3.00})
	shoppingList = append(shoppingList, Item{category: 0, name: "Cups", quantity: 5, unitCost: 3.00})
	shoppingList = append(shoppingList, Item{category: 1, name: "Bread", quantity: 2, unitCost: 2.00})
	shoppingList = append(shoppingList, Item{category: 1, name: "Cake", quantity: 3, unitCost: 1.00})
	shoppingList = append(shoppingList, Item{category: 2, name: "Coke", quantity: 5, unitCost: 2.00})
	shoppingList = append(shoppingList, Item{category: 2, name: "Sprite", quantity: 5, unitCost: 2.00})

}

func getShoppingList() []Item {
	return shoppingList
}

func showShoppingList() {
	for _, item := range shoppingList {
		fmt.Printf("Category: %d - ", item.category)
		fmt.Printf("Item: %s ", item.name)
		fmt.Printf("Quantity: %d ", item.quantity)
		fmt.Printf("Unit Cost: %f\n", item.unitCost)
	}
}

/*
	function to add an item to the Shopping List
*/
func addItem(i Item) {
	shoppingList = append(shoppingList, i)
}

/*
	function to update an item in the Shopping List
*/
func updateItem(updatedItem Item) (bool, error) {

	// find the item
	for _, item := range shoppingList {
		if item.name == updatedItem.name {
			item.name = updatedItem.name
			item.category = updatedItem.category
			item.quantity = updatedItem.quantity
			item.unitCost = updatedItem.unitCost
			break
		}
	}

	return true, nil
}

/*
	function to find the first matching element (by name)
	and return the element index.
	if no matching element is found, return -1.
*/
func findFirstMatch(list *[]Item, name string) int {
	for i, element := range *list {

		if element.name == name {
			return i
		}
	}

	return -1
}
