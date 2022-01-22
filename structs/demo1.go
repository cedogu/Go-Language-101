package structs

import "fmt"

//structs in go = classes in C#

func Demo1() {
	fmt.Println(product{"Computer", 400, "Lenovo"})
	fmt.Println(product{"Computer", 800, "Apple"})
	fmt.Println(product{name: "Computer", unitPrice: 499})

}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
