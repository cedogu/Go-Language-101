package conditionals

import "fmt"

func Demo() {
	var account float64 = 1000
	var withdrawnMoney float64 = 900

	if withdrawnMoney > account {
		fmt.Println("insufficient balance")

	}

	if withdrawnMoney <= account {
		fmt.Println("Your amount is getting ready")
		account = account - withdrawnMoney
	}
	fmt.Println("Finish. Amount : " + fmt.Sprint(account)) //%v = default value, we don't need to write this. it comes automatically.

}
