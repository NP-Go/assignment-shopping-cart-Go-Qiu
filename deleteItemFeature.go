package main

import "fmt"

func showDeleteItemScreen() {

	name := ""

	clearScreen()

	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("+ Delete Item                          +")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")

	for {
		fmt.Println("Enter name of item to be deleted: ")
		fmt.Scanln(&name)
		if isEmpty, message := inputNotEmptyCheck(&name, "Name"); !isEmpty {
			fmt.Println(message)
		} else {
			// passed.
			// check if item existed.
			var keys []string
			for k, _ := range shoppingCart {
				keys = append(keys, k)
			}

			if !containsString(keys, name) {
				// not found
				fmt.Println("")
				fmt.Println("Item Not Found !")
			} else {
				// found
				deleteItem(name)
				fmt.Println("")
				fmt.Printf("Deleted %s from the Shopping List.", name)
			}
			break
		}
	}
}
