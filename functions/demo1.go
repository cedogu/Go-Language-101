package functions

import "fmt"

func Math(nr1 int, nr2 int) int {
	var result = nr1 + nr2
	return result
}

func SayHi(userName string) {
	fmt.Println("Hello,", userName)
}
