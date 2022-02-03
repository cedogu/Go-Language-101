package interfaces

import "fmt"

func dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)

	//buradaki ok'un görevi ok değer int ise true, değilse false veriyor. aşağıda "doruk" yazdığım için
	//sonucunu da aşağıda olduğu gibi false alıyoruz.

}

func Demo3() {
	var sayi1 interface{} = "Doruk"
	dogrula(sayi1)

	var sayi2 interface{} = 5
	dogrula(sayi2)
}
