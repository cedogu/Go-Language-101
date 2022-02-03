package errorhandling

import (
	"fmt"
)

type borderException struct {
	message   string
	parameter int
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func GuessNr(guess int) (string, error) {
	if guess < 1 || guess > 100 {
		return "", &borderException{parameter: 102, message: "out of borders"}
	}
	return "Bildiniz", nil

}
