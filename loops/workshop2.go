package loops

import "fmt"

func Workshop2() {

	number := 0
	fmt.Println("Please give a number")
	fmt.Scanln(&number)

	primenumber := true

	for i := 2; i < number; i++ {
		if number%i == 0 {
			primenumber = false
		}
	}
	if primenumber == true {
		fmt.Println("Your result is a prime number")
	} else {
		fmt.Println("Your result is a not prime number")
	}

}
