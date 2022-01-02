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

		//i'yi 0'dan baslattık ve döngü icindeki tüm sayıları dolasacak.
		//i'de bulunan hangi sayı en büyük sayı bizim numbers'ın i değerine eşit olacak.
		//i ilk önce 10 ile baslayacak, sonra 20-30-40 ve en son 50'de kalacak cünkü en büyük sayımız 50'ye eşit olacak.

	}

	fmt.Println("Biggest number is : ", biggest)
}
