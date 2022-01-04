package for_range

import "fmt"

func Demo3() {

	dictionary := map[string]string{"Cam": "Glass", "Muz": "Banana", "Kedi": "Cat"}

	for k, v := range dictionary {
		fmt.Print(k)
		fmt.Print(v)
	}

}
