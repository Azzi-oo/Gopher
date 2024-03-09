package main

import "fmt"

func main() {
	fmt.Println(findMin(1, 2, 3, 5434, -345, 53, 432))
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]

	for _, i := range numbers {
		if i > max {
			max = i
		}
	}
	return max
}
