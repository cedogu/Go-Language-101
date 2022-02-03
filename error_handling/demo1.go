package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt")
	//nil
	if err != nil {
		fmt.Println("File can not be found")
	} else {
		fmt.Println(f.Name())
	}
	//eğer dosyayı bulamazsa nil, bulursa f(file) 'ı yazdırıyoruz.

	//when we drag the demo1.txt file to the main area, there won't be any error.
	//but if we don't drag it and keep the demo1.txt file in the error_handling folder, we will catch the error
}
