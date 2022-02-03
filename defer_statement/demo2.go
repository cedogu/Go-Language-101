package deferstatement

import "fmt"

func Demo2(number int32) string {
	defer DeferFunc()

	if number%2 == 0 {
		fmt.Println("Even number in progress")
		return "Even number"
	}
	if number%2 != 0 {
		fmt.Println("Odd number in progress")
		return "Odd number"
	}

	return "Not clear"

}

func Test() {
	result := Demo2(10)
	fmt.Println(result)

}

func DeferFunc() {
	fmt.Println("Over")
}
