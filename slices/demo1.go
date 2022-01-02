package slices

import "fmt"

func Demo1() {
	names := make([]string, 3)

	fmt.Println(names)
	names[0] = "Cem"
	names[1] = "Doruk"
	names[2] = "Gürsoy"

	names = append(names, "Ferahnaz")

	fmt.Println(names)
	fmt.Println(len(names))

	//dizilere yeni elemanı append ile ekliyoruz

}
