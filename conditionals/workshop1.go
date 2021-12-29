package conditionals

import "fmt"

func Workshop() {

	var n1, n2, n3 int = 100, 5, 28

	var biggest int = n1

	if n2 > biggest {
		biggest = n2
	}

	if n3 > biggest {
		biggest = n3
	}

	fmt.Printf("biggest number is: %v", biggest)
	//we can also write as fmt.Print("biggest number is :",biggest)

}
