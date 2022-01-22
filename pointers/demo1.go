package pointers

import "fmt"

func Demo1(number *int) {
	*number = *number + 1
	fmt.Println("number in the demo func is:", *number)

	//değer tiplilerdekiler steak'te olur ama ref tipliler heap'te olur.
	//bool, int vb değer tiplilerdir, yani demo'daki değer ve main'deki değer farklı farklı olusumlardır ve birbirlerinin yeirni almazlar
	//yani bellekte iki farklı int number değerinde yer açılıyor
	//array, class vb. ref tiplidir.
	//go pointers == c# reference type
}
