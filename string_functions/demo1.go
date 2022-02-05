package stringfunctions

import (
	"fmt"
	"strings"
	s "strings"
)

func Demo1() {
	name := "Doruk"
	fmt.Println(strings.Count(name, "d"))
	fmt.Println(strings.Contains(name, "D"))
	fmt.Println(strings.Index(name, "D"))

	result := s.Index(name, "d")

	if result != -1 {
		fmt.Println("There is letter a")
	} else {
		fmt.Println("There isn't any letter a")
	}

	fmt.Println(s.ToLower(name))
	fmt.Println(s.ToUpper(name))

	//case sensitive. D yazdığımda sonuc 1 iken d yazınca sonuc 0! Büyük küçük ayrımı!
	//ascii
}
