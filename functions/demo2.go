package functions

func FourOperations(nr1 int, nr2 int) (int, int, int, float32) {

	addition := nr1 + nr2
	substraction := nr1 - nr2
	multiplication := nr1 * nr2
	division := nr1 / nr2

	return addition, substraction, multiplication, float32(division)

	//type return has been transferred automatically to float32 by go. cast operation.

}

//we defined the return type of our operations.
// tüm matematiksel işlemlerini döndüren bir fonksiyon olacak.
