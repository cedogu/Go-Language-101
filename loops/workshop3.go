package loops

import "fmt"

func Workshop3() {

	number1 := 220
	number2 := 284
	result1 := 0
	result2 := 0

	for i := 1; i < number1; i++ {
		if number1%i == 0 {

			result1 = result1 + i
		}
	}
	for i := 1; i < number2; i++ {
		if number2%i == 0 {
			result2 = result2 + i
		}
	}

	if result1 == number2 && result2 == number1 {
		fmt.Printf("%v and %v are Friendly Numbers", number1, number2)
	} else {
		fmt.Println("Your result is not Friendly Numbers")
	}

}
