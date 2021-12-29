package main

import "fmt"

func main() {

	/*var metin string = "Hi Turkey"

	fmt.Println(metin)
	fmt.Println("Hello, World! ") */

	// var kdv int = 10
	// fmt.Println(kdv)
	// fmt.Println(100 + 100*10/100)
	// fmt.Println()

	// var kdv2 float32 = 0.1
	// fmt.Println(kdv2)
	// fmt.Println(100 + 10*10/kdv2)

	kdv3 := 25 //different definition in Go language. I've never seen this before. we can write anything we want here. The programme will define its type automatically.
	fmt.Println(kdv3)
	//fmt.Printf("data type %T", kdv3)  with this definition we can learn the type of our defition.

	name := "doruk" //this is another example.
	fmt.Println(name)

	var durum bool
	var metin1 string = "Cem Doruk"
	var metin2 string = "Savaş"

	durum = metin1 == metin2
	//durum = metin1 != metin2 if we wrote so, the result would be true cuz both names are different and not equal.

	fmt.Println(durum)
	fmt.Println(!durum)

}
