package deferstatement

import "fmt"

func A() {
	fmt.Println("A function is working")
}
func C() {
	fmt.Println("C function is working")
}
func D() {
	fmt.Println("D function is working")
}
func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("B function is working")

}

//Defer- last in first out
