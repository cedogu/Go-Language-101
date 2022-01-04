package maps

import "fmt"

func Demo1() {
	//key value
	//we can use "make" functions to create a map.
	dictionary := make(map[string]string)
	dictionary["table"] = "Masa"
	dictionary["book"] = "Kitap"
	dictionary["cat"] = "Kedi"
	dictionary["pencil"] = "Kalem"
	dictionary["computer"] = "Bilgisayar"
	//yazarken bir şeyi eşleştirirsek ona bağlıyor, vermezsek de default değerini veriyo

	fmt.Println(dictionary["book"])
	fmt.Println("Element Number :", len(dictionary))
	fmt.Println("Dictionary:", dictionary)
	delete(dictionary, "table")
	fmt.Println("Element Number:", len(dictionary))
	fmt.Println("Dictionary:", dictionary)

	value, existence := dictionary["table"] //if we write something irrelevant it becomes false automatically,but related things will be true.
	fmt.Println(value)                      //table will also become false cuz above we deleted with the delete function.
	fmt.Println("Being on the list:", existence)

	//dictionary de arka planda dizi yapısına sahiptir.

	dictionary2 := map[string]string{"Glass": "bardak", "Phone": "Telefon", "Microphone": "Mikrofon"}
	fmt.Println(dictionary2) //another way to write a map in GO

}
