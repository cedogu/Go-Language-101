package loops

import "fmt"

func Workshop1() {
	nrinmymind := 80
	guessednr := 0
	gameright := 0

	fmt.Println("Give a number:")
	fmt.Scanln(&guessednr)
	fmt.Println(guessednr)
	gameright = gameright + 1

	for nrinmymind != guessednr {
		if nrinmymind < guessednr {
			fmt.Println("Wrong, maybe lower?")
			fmt.Scanln(&guessednr)
			gameright = gameright + 1
		}
		if nrinmymind > guessednr {
			fmt.Println("Wrong, maybe higher?")
			fmt.Scanln(&guessednr)
			gameright = gameright + 1
		}

	}

	successrate := " "
	if gameright > 0 && gameright <= 3 {
		successrate = "Super!"
	} else if gameright <= 6 {
		successrate = "Good!"
	} else {
		successrate = "Not Bad"
	}
	fmt.Printf("Congrats!, you win in the %v right : %v", gameright, successrate)

}
