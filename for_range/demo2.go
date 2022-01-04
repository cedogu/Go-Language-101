package for_range

import "fmt"

func Demo2() {

	//Odd numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := 0

	for _, number := range numbers {

		if number%2 != 0 {
			result = result + number
		}
	}
	fmt.Println("Result:", result)

	//even numbers
	numbers2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result2 := 0

	for _, number2 := range numbers2 {
		if number2%2 == 0 {
			result2 = result2 + number2

		}
	}
	fmt.Println("Result", result2)
}
