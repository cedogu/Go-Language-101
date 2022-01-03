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

	functions.SayHi("Doruk")

	var result = functions.Math(2, 5)
	fmt.Println(result * 10) //we can make a quick calculation with some little tricks.

}
