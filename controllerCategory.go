package main

import (
	"sort"
)

// type intKeyValuePair struct {
// 	key   int
// 	value string
// }

var mCategory map[int]string

/*
	function to pre-load the default categroy map
*/
func preloadCategoryMap() {
	mCategory = make(map[int]string)

	// Initialize the category map
	mCategory[0] = "Household"
	mCategory[1] = "Food"
	mCategory[2] = "Drinks"

}

/*
	function to get the category map
	listed by the key, in asc order
*/
func listCategories() []intKeyValuePair {
	// declare the category map

	var keys []int
	var sCategory []intKeyValuePair

	for k := range mCategory {
		keys = append(keys, k)
	}

	// sort the keys slice in asc order
	sort.Ints(keys)
	for _, k := range keys {
		sCategory = append(sCategory, intKeyValuePair{key: k, value: mCategory[k]})
	}

	return sCategory
}

/*
	function to add a category into the map
*/
func addCategory(key int, category string) {

}
