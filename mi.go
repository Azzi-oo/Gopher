package main

import "fmt"

// var msg string

func main() {
	number := 5
	var p *int

	p = &number

	fmt.Println(p)
	fmt.Println(&number)

	// message := "Scoro i be ninja"
	// fmt.Println(&message)
	// changeMessage(message)
	// fmt.Println(&message)
}

func changeMessage(message string) {
	// *message += " (from func printMessage())"
	fmt.Println(&message)
}
