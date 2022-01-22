package pointers

import "fmt"

func Demo2(numbers []int) {
	numbers[0] = 100

	fmt.Println(numbers[0])
}
