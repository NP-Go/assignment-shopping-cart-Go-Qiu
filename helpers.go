package main

import (
	"bufio"
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

/*
	function to clear the the screen (macOS, windows or linux)
*/

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

/*
	function to check if the input value (string) is valid and has a matcing key with respect
	to the keys (slice) in the lookup data.
	return true (and message) when input value is valid and has a match
	return false (and message) when input value is not valid or has not got a match
*/
func inputIsValidLookupValueCheck(input *string, inputName string, lookupData *[]intKeyValuePair) (bool, string) {
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
	function to get all the values (in a slice) of a specific lookup data slice
*/
func getLookupValues(lookupData []intKeyValuePair) []string {
	var values []string
	for _, c := range lookupData {
		values = append(values, c.value)
	}
	return values
}

/*
	function to execute check on the specific input value
	to ensure it is a valid numeric float value
*/
func inputIsValidFloatValueCheck(input *string, requiredPositive bool, inputName string) (bool, string) {

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

/*
	function to check if the input value is not an empty string
	return true (and message) if it is not an empty string
	return false (and message) if it is an empty string
*/
func inputNotEmptyCheck(input *string, inputName string) (bool, string) {
	var tag string
	if len(strings.TrimSpace(inputName)) != 0 {
		tag = strings.TrimSpace(inputName)
	} else {
		tag = "Input value"
	}
	if len(strings.TrimSpace(*input)) == 0 && strings.TrimSpace(inputName) != "-" {
		// empty string
		return false, tag + " cannot be empty."
	} else if len(strings.TrimSpace(*input)) == 0 && strings.TrimSpace(inputName) == "-" {
		// empty string but use 'No Input Found' message
		return false, ""
	} else {
		// non-empty string
		// exit from infinite loop

		return true, "Ok"
	}
}

/*
	function to check if the input value (string) is unique
	with respect to the existing list (slice of values)
	return true (and message) if it is unique
	return false (and message) if it is not unique
*/
func inputIsUniqueByNameCheck(input *string, s []string) (bool, string) {

	message := ""
	var isUnique bool

	for i, element := range s {
		if element == strings.TrimSpace(*input) {
			// found, not unique
			message = "Category : " + *input + " already exist at " + strconv.Itoa(i)
			isUnique = false
			break
		} else {
			// not found, is unique
			isUnique = true
		}
	}
	return isUnique, message
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

func containsString(s []string, v string) bool {
	found := false
	for _, element := range s {
		if element == v {
			// found
			found = true
			break
		}
	}
	return found
}

/*
** function to get an input (string) that has space between word
 */
func getInputString() string {
	inputString := bufio.NewReader(os.Stdin)
	line, err := inputString.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(line, "\n")
}
