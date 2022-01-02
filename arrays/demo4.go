package arrays

import "fmt"

func Demo4() {
	var numbers [2][3]int

	numbers[0][0] = 1
	numbers[0][1] = 3
	numbers[0][2] = 5
	numbers[1][0] = 2
	numbers[1][1] = 4
	numbers[1][2] = 6
	//2 satır 3 ise sütun sayısıdır. yani bir array 2 satırlı ve üc sutunlu
	//0-1-2 sayma sayısı ile diziler başlar. bu yüzden 0-1-2 olarak görülür

	for i := 0; i < 2; i++ { //satır sayısı 2 sütun sayısı 3
		for j := 0; j < 3; j++ {
			fmt.Print(numbers[i][j])
			fmt.Print(" ")
		}
		fmt.Println()

	} //2 means numbers of lines

	//fmt.Println(numbers[1][1])
}
