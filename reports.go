package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CategoryTotalCost struct {
	category  int
	totalCost float64
}

type Item struct {
	category int
	name     string
	quantity int
	unitCost float64
}

type ItemWithCategory struct {
	category string
	name     string
	quantity string
	unitCost float64
}

/*
	function to calculate and print the
	total cost of items (group by category)
*/
func getTotalCostByCategory(m map[string]ItemInfo) map[int]float64 {

	// cateogyr total Cost slice
	var tc map[int]float64
	tc = make(map[int]float64)

	for _, uc := range m {

		tc[uc.category] += uc.unitCost * float64(uc.quantity)
	}
	return tc
}

func getItemsByCategory(m map[string]ItemInfo) []string {

	var mCat map[int]string
	mCat = make(map[int]string)

	var itemsList []string

	// loop through the shopping cart map
	// to create a map of categories (with associated Item data as string)
	for key, info := range m {

		categoryName := getCategoryByKey(info.category)
		itemAsString := "Category : " + categoryName + " - "
		itemAsString += "Item : " + key + " "
		itemAsString += "Quantity : " + strconv.Itoa(info.quantity) + " "
		itemAsString += "Unit Cost : $ " + strconv.FormatFloat(info.unitCost, 'f', 2, 64)

		if len(strings.TrimSpace(mCat[info.category])) == 0 {
			mCat[info.category] += itemAsString
		} else {
			itemAsString = ";" + itemAsString
			mCat[info.category] += itemAsString
		}
	}

	// loop through the map of categories
	// to build a slice of Items sorted by categories
	for _, items := range mCat {

		i := strings.Split(items, ";")
		itemsList = append(itemsList, i...)
	}
	return itemsList
}

func printTotalCostByCategory(m map[string]ItemInfo) {

	clearScreen()

	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("+ Report - Total Cost by Category      +")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")

	mCatTotalCost := getTotalCostByCategory(m)
	for key, tc := range mCatTotalCost {
		fmt.Printf("%s : $ ", getCategoryByKey(key))
		fmt.Printf("%0.2f \n", tc)
	}
	fmt.Println("")

	returnToPrevious("REPORTS")
}

func printItemsByCategory(m map[string]ItemInfo) {

	clearScreen()

	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("+ Report - List of Items by Category   +")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")
	itemsList := getItemsByCategory(shoppingCart)
	for _, item := range itemsList {
		fmt.Println(item)
	}
	fmt.Println("")

	returnToPrevious("REPORTS")

}
