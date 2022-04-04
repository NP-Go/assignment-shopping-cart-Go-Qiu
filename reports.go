package main

import "fmt"

type CategoryTotalCost struct {
	category  int
	totalCost float64
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

func printTotalCostByCategory(m map[string]ItemInfo) {

	clearScreen()

	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("+ Report - Total Cost by Category      +")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")

	// var mItemCatUnitCost map[int]float64
	// mItemCatUnitCost = make(map[int]float64)

	// for _, v := range m {
	// 	mItemCatUnitCost[v.category] += (v.unitCost * float64(v.quantity))
	// }

	mCatTotalCost := getTotalCostByCategory(m)
	for key, tc := range mCatTotalCost {
		fmt.Printf("%s : $ ", getCategoryByKey(key))
		fmt.Printf("%0.2f \n", tc)
	}
	fmt.Println("")

}
