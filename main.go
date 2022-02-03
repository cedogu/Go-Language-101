package main

import (
	"fmt"
	errorhandling "golesson/error_handling"
)

func main() {

	// variables.Demo1()
	// fmt.Print()
	// conditionals.Demo2()
	// conditionals.Workshop()
	// loops.Workshop3()
	//arrays.Demo4()
	//slices.Demo2()

	// functions.SayHi("Doruk")

	// var result = functions.Math(2, 5)
	// fmt.Println(result * 10) //we can make a quick calculation with some little tricks.

	// result1, result2, result3, result4 := functions.FourOperations(10, 6)
	// fmt.Println(result1, result2, result3, result4)
	// fmt.Println("Addition: ", result1)
	// fmt.Println("subtraction: ", result2)
	// fmt.Println("multiplication: ", result3)
	// fmt.Println("division: ", result4)

	// var sonuc = functions.ToplaVariadic(1, 4, 6, 3, 10)
	// fmt.Println(sonuc)
	// fmt.Println(functions.ToplaVariadic(10, 12, 3))
	// fmt.Println(functions.ToplaVariadic(2, 5))

	// //iki türlü yazma şekli vardır. biz yukarıdaki gibi veirriz ve Go otomatik olarak onları DİZİ yapar.
	// //ya da aşagıdaki gibi direkt olarak array verebiliriz.
	// //elimizdeki array'ı da buraya yollayabiliriz.
	// sayilar := []int{4, 5, 6}
	// fmt.Println(functions.ToplaVariadic(sayilar...))

	// maps.Demo1()

	//for_range.Demo1()
	//for_range.Demo2()
	//for_range.Demo3()

	// number := 20
	// pointers.Demo1(&number)
	// fmt.Println("Number in the main func is: ", number)

	// numbers := []int{1, 2, 3}
	// pointers.Demo2(numbers)
	// fmt.Println(numbers[0])

	// structs.Demo2()

	// go goroutines.EvenNumbers() //to create a routine we just add "go", pretty easy!
	// go goroutines.OddNumbers()
	// time.Sleep(1 * time.Second)
	// fmt.Println("Main is over")

	// evenNumbersCn := make(chan int)
	// oddNumbersCn := make(chan int)
	// go channels.EvenNumbers(evenNumbersCn)
	// go channels.OddNumbers(oddNumbersCn)

	// evenNumbersAdd, oddNumbersAdd := <-evenNumbersCn, oddNumbersCn

	// multiplying := evenNumbersAdd * <-oddNumbersAdd
	// fmt.Println("multiplying: ", multiplying)

	// interfaces.Demo1()

	// interfaces.Demo2()
	// interfaces.Demo3()

	// deferstatement.B()
	// deferstatement.Test()
	// deferstatement.Demo3()

	// errorhandling.Demo1()

	fmt.Println(errorhandling.GuessNr(102))

}
