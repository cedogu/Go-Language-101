package interfaces

import (
	"fmt"
	"math"
)

//interface metot imzaları barındırır

type shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape) { //--buradaki shape asagıdaki rectangle'nın bellekteki yerini tutuyor.
	fmt.Println("Area : ", s.area())
	//bu imza sonrası içinde area olan her şeyi asagıda demo1'e yollayabiliriz.
}

func Demo1() {
	r := rectangle{width: 10, height: 6}
	school(r)

	c := circle{radius: 10}
	school(c)
}
