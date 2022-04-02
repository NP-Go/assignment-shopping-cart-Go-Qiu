package main

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

/*
	function to check on the os
	where the app is running in.
*/
func clearScreen() {
	os := runtime.GOOS
	switch os {
	case "windows":
		clearStdOut("cls")
	case "darwin":
		clearStdOut("clear")
	default:
		clearStdOut("clear")
	}
}

func clearStdOut(command string) {
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

/*
	function to execute checks on the specific input
	to ensure it is a valid numeric integer value
*/
func inputIsValidIntegerValue(input *string, requiredPositive bool, inputName string) (bool, string) {
	var tag string
	entry, err := strconv.Atoi(*input)

	if len(strings.TrimSpace(inputName)) != 0 {
		tag = strings.TrimSpace(inputName)
	} else {
		tag = "Input value"
	}
	// generic rules
	if err != nil {
		// input cannot be converted to integer
		return false, tag + " must be an integer value."
	} else if entry == 0 {
		return false, tag + " must not be zero."
	} else {
		// passed the generic rules

		if requiredPositive {

			// check for positive integer value
			if entry < 0 {
				return false, tag + " must be a postive number."
			} else {
				// passed generic rules and positive integer value check
				return true, "Ok"
			}
		} else {
			// passed generic checks only
			return true, "ok"
		}
	}
}

func inputIsValidLookupValue(input *string, inputName string, lookupData *[]intKeyValuePair) (bool, string) {
	var tag string
	entry, err := strconv.Atoi(*input)

	if len(strings.TrimSpace(inputName)) != 0 {
		tag = strings.TrimSpace(inputName)
	} else {
		tag = "Input value"
	}
	// generic rules
	if err != nil {
		// input cannot be converted to integer
		return false, tag + " must be an integer value."
	} else {
		// passed the generic rules

		// get all the keys in the category lookup
		keys := getLookupKeys(*lookupData)
		if !containsInt(keys, entry) {
			return false, tag + " must be an integer value in the Category Lookup."
		} else {
			// passed lookup key check
			return true, "Ok"
		}
		//
	}
}

/*
	function to get all the keys (in a slice) of a specific lookup data slice
*/
func getLookupKeys(lookupData []intKeyValuePair) []int {
	var keys []int
	for _, c := range lookupData {
		keys = append(keys, c.key)
	}
	return keys
}

/*
	function to execute check on the specific input value
	to ensure it is a valid numeric float value
*/
func inputIsValidFloatValue(input *string, requiredPositive bool, inputName string) (bool, string) {

	var tag string
	entry, err := strconv.ParseFloat(*input, 32)

	if len(strings.TrimSpace(inputName)) != 0 {
		tag = strings.TrimSpace(inputName)
	} else {
		tag = "Input value"
	}
	// generic rules
	if err != nil {
		// input cannot be converted to integer
		return false, tag + " must be a numeric value."
	} else if entry == 0 {
		return false, tag + " must not be zero."
	} else {
		// passed the generic rules

		if requiredPositive {

			// check for positive integer value
			if entry < 0 {
				return false, tag + " must be a postive number."
			} else {
				// passed generic rules and positive integer value check
				return true, "Ok"
			}
		} else {
			// passed generic checks only
			return true, "ok"
		}
	}
}

func inputNotEmptyCheck(input *string, inputName string) (bool, string) {
	var tag string
	if len(strings.TrimSpace(inputName)) != 0 {
		tag = strings.TrimSpace(inputName)
	} else {
		tag = "Input value"
	}
	if len(strings.TrimSpace(*input)) == 0 {
		// empty string
		return false, tag + " cannot be empty."
	} else {
		// non-empty string
		// exit from infinite loop

		return true, "Ok"
	}
}

/*
	function to check if a specific slice of int values
	contains a specific int value
*/
func containsInt(s []int, v int) bool {
	found := false
	for _, element := range s {
		if element == v {
			//found
			found = true
			break
		}
	}
	return found
}
