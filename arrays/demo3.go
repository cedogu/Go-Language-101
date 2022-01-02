package arrays

import "fmt"

func Demo3() {
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println(numbers)

	biggest := numbers[0]

	//instead of writing sayilar.length we use only len (sayilar) in GO

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > biggest {
			biggest = numbers[i]
		}

	}

	fmt.Println("Biggest number is : ", biggest)
}
