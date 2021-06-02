package main

import (
	"fmt"
)

func main() {

	var stringVal string = "(The{brackets[are]helpful}for)programming"
	fmt.Println(stringVal[0:1])
	fmt.Println(len(stringVal))

	for i := 0; i < len(stringVal); i++ {

	}
}
