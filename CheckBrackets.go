package main

import "fmt"

func main() {

	// hash map to check for open brackets
	var openBrackets map[string]bool = map[string]bool{
		"(": true,
		"{": true,
		"[": true,
	}

	// hash map to match the open brackets with the closed bracket
	var brackets map[string]string = map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	// Input string 1
	var stringVal1 string = "(The{brackets[are]helpful}for)programming"
	fmt.Println("Input 1: " + stringVal1)
	fmt.Println(checkForBalancedBrackets(stringVal1, openBrackets, []string{}, brackets))

	// Input string 2
	var stringVal2 string = "{There(is{a[dangling]bracket}in)front"
	fmt.Println("Input 2: " + stringVal2)
	fmt.Println(checkForBalancedBrackets(stringVal2, openBrackets, []string{}, brackets))

	// Input string 3
	var stringVal3 string = "There(is{a[dangling]bracket}in)the{end"
	fmt.Println("Input 3: " + stringVal3)
	fmt.Println(checkForBalancedBrackets(stringVal3, openBrackets, []string{}, brackets))

}

// stringVal: input string
// openBrackets: containes the list of open brackets
// stack: stack to keep track of the open brackets encountered
// brackets: matching closing brackets with opening brackets

// As we traverse the entire characters in the input string the
// compalixity of the algorithm is O(n) where n is the total
// length of the input string
func checkForBalancedBrackets(stringVal string, openBrackets map[string]bool, stack []string, brackets map[string]string) (result string) {

	// loop through the input string's indivisual characters
	for i := 0; i < len(stringVal); i++ {

		// get the ith character
		var tempChar string = stringVal[i : i+1]

		// if the char is an open bracket
		if openBrackets[tempChar] {
			stack = append(stack, tempChar)
		} else {

			//check if the bracket is a closed bracket
			openBracketCheck, ok := brackets[tempChar]
			if ok {

				// if the length the stack is 0, that means we did not
				// encounter open bracket for the current closed bracket
				if len(stack) == 0 {
					return "NO"
				}

				// pop the top of the stack
				var popedChar string = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				// check if the closeing bracket that corresponds to
				// poped element of the stack matches the
				// closed bracket that we have encountered currently
				if openBracketCheck != popedChar {
					return "NO"
				}
			}
		}

	}

	// check if there are any dangling open brackets in the stack
	// which did not get a matching closing bracket in the loop
	if len(stack) == 0 {
		return "YES"
	} else {
		return "NO"
	}

}
