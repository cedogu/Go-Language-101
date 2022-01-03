package main

import (
	"fmt"
	"golesson/functions"
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

	result1, result2, result3, result4 := functions.FourOperations(10, 6)
	fmt.Println(result1, result2, result3, result4)
	fmt.Println("Addition: ", result1)
	fmt.Println("subtraction: ", result2)
	fmt.Println("multiplication: ", result3)
	fmt.Println("division: ", result4)

}
