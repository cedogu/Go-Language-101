package channels

import (
	"fmt"
	"time"
)

func EvenNumbers(evenNumbersCn chan int) {
	addition := 0
	for i := 0; i <= 10; i += 2 {
		addition = addition + i
		fmt.Println("Even number in progress")
		time.Sleep(1 * time.Second)
	}

	evenNumbersCn <- addition
}

func OddNumbers(oddNumbersCn chan int) {
	addition := 0
	for i := 1; i <= 10; i += 2 {

		addition = addition + i
		fmt.Println("Odd number is in progress")
		time.Sleep(1 * time.Second)
	}
	oddNumbersCn <- addition
}

//fonksiyon içine parametre olarak channel yolluyoruz, Channel ismi - channel ve channel türünü yolluyoruz.
