package conditionals

import "fmt"

func Demo2() {
	var account float64 = 1000
	var withdrawnMoney float64 = 1000

	if withdrawnMoney > account {
		fmt.Println("insufficient balance")

	} else if withdrawnMoney == account {
		fmt.Println("Your money is getting ready")
		fmt.Println("No money left in account")
	} else {
		fmt.Println("")
	}

}
