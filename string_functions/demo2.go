package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	name := "Doruk"
	fmt.Println(s.HasPrefix(name, "D"))
	fmt.Println(s.HasSuffix(name, "k"))
	fmt.Println(s.Index(name, "do"))

	letters := []string{"D", "O", "R", "U", "K"}
	result := s.Join(letters, "*")
	fmt.Println(result)

	fmt.Println(s.Replace(result, "*", "+", 3))
	//-1 ne görürsen degistir demek, 0 hiç değişme demek
	//pozitif sayı kaç verirsek o kadar degisim olur.
	fmt.Println(s.Split(name, "*"))
	fmt.Println(s.Repeat(name, 5))

}
