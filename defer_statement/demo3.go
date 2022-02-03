package deferstatement

import "fmt"

type Product struct {
	productName string
	unitPrice   int
}

func (p Product) save() {
	fmt.Println("Product saved :", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Logged")
}

func Demo3() {

	p := Product{productName: "Laptop", unitPrice: 5000}

	p = Product{productName: "Mouse", unitPrice: 40}
	fmt.Println("Over")
	fmt.Println(p.productName)
	defer p.save()
	//defer'i Mouse'ın altına koysak sonuc mouse olur, laptop altında olsa sonuc laptop olur.
	//sebebi ise defer fonksiyonu en neyde bırakırsa onu baz alır.
	//genelde en sona konulması daha iyi olur.
}
