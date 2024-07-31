package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Gophers! Pass me a word and I will print the last character!")

	args := os.Args[1]
	fmt.Println("The last character in that word is: %c\n", args[len(args)])
}
