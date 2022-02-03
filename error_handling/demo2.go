package errorhandling

import (
	"errors"
	"fmt"
)

func GuessNumber(guess int) (string, error) {

	numberInMyMind := 50

	if guess < 1 || guess > 100 {
		return "", errors.New("please give a number betweeen 1-100")
	}

	if guess > numberInMyMind {
		return "Please give a smaller number", nil
	}

	if guess < numberInMyMind {
		return "please give a greater number", nil
	}

	return "You win!", nil
}

func Demo2() {

	message, _ := (GuessNumber(50))
	fmt.Println(message)

}
