package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(prediction("ср"))
}

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "пн":
		return "Хорошего начала азны", nil
	case "вт":
		return "Good you vtornik", nil
	case "ср":
		return "Good avg of week!", nil
	default:
		return "ne valid day of week", errors.New("invalid day of the week")
	}
}
