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

type ItemInfo struct {
	category int
	quantity int
	unitCost float64
}

// declare a slice to cache the items data
// var shoppingList = []Item{}
var shoppingCart map[string]ItemInfo

/*
	function to preload the data in the shopping list
*/
func preloadShoppingList() {

	// shoppingList = append(shoppingList, Item{category: 0, name: "Fork", quantity: 4, unitCost: 3.00})
	// shoppingList = append(shoppingList, Item{category: 0, name: "Plates", quantity: 4, unitCost: 3.00})
	// shoppingList = append(shoppingList, Item{category: 0, name: "Cups", quantity: 5, unitCost: 3.00})
	// shoppingList = append(shoppingList, Item{category: 1, name: "Bread", quantity: 2, unitCost: 2.00})
	// shoppingList = append(shoppingList, Item{category: 1, name: "Cake", quantity: 3, unitCost: 1.00})
	// shoppingList = append(shoppingList, Item{category: 2, name: "Coke", quantity: 5, unitCost: 2.00})
	// shoppingList = append(shoppingList, Item{category: 2, name: "Sprite", quantity: 5, unitCost: 2.00})

	shoppingCart = make(map[string]ItemInfo)
	shoppingCart["Fork"] = ItemInfo{category: 0, quantity: 4, unitCost: 3.00}
	shoppingCart["Plates"] = ItemInfo{category: 0, quantity: 4, unitCost: 3.00}
	shoppingCart["Cups"] = ItemInfo{category: 0, quantity: 5, unitCost: 3.00}
	shoppingCart["Bread"] = ItemInfo{category: 1, quantity: 2, unitCost: 2.00}
	shoppingCart["Cake"] = ItemInfo{category: 1, quantity: 3, unitCost: 1.00}
	shoppingCart["Coke"] = ItemInfo{category: 2, quantity: 5, unitCost: 2.00}
	shoppingCart["Sprite"] = ItemInfo{category: 2, quantity: 5, unitCost: 2.00}

}

func getShoppingList() map[string]ItemInfo {
	return shoppingCart
}

func showShoppingList() {

	for key, item := range shoppingCart {
		fmt.Printf("Category: %s - ", getCategoryByKey(item.category))
		fmt.Printf("Item: %s ", key)
		fmt.Printf("Quantity: %d ", item.quantity)
		fmt.Printf("Unit Cost: %f\n", item.unitCost)
	}
}

/*
	function to add an item to the Shopping List
*/
func addShoppingListItem(k string, i ItemInfo) {
	shoppingCart[k] = i
}

/*
	function to update an item in the Shopping List
*/
func updateShoppingListItem(k string, u ItemInfo) (bool, error) {

	shoppingCart[k] = u
	return true, nil
}

/*
	function to delete an item from the shopping list
*/
func deleteItem(k string) {
	delete(shoppingCart, k)
}

/*
	function to print the current data set in the shopping cart
*/
func printShoppingListData() {

	for k, v := range shoppingCart {
		fmt.Printf("%s", k)
		fmt.Printf(" - %v \n", v)
	}

}
