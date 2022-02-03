package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo12.txt") //I changed the name by writing demo12. (main demo1 is in the main area)
	//nil
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("File could not be found", pErr.Path)
			return
		} else {
			fmt.Println("File can not be found")
			return
		}

	} else {
		fmt.Println(f.Name())
	}
	//eğer dosyayı bulamazsa nil, bulursa f(file) 'ı yazdırıyoruz.

	//when we drag the demo1.txt file to the main area, there won't be any error.
	//but if we don't drag it and keep the demo1.txt file in the error_handling folder, we will catch the error
}
