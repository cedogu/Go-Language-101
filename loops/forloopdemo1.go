package loops

import "fmt"

func Demo1() {
	var text string = "Hi, World"

	//first create an index as "i" and then give it a value such as "1". then build your loop and inside the loop add your string.

	i := 1

	for i <= 5 {
		fmt.Println(text)
		i = i + 1
	}
	fmt.Print("Loop is over, thanks for watching.")
}
