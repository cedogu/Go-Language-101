package functions

//sayilari dizi olarak ele alıyor

func ToplaVariadic(sayilar ...int) int {
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam = toplam + sayilar[i]

	}
	return toplam
}

//birden fazla sayi verebileceğimizi belirttigimiz icin "sayilar" dedik.
//... sonsuz sayı verebileceğimizi belirtiyoruz, yani eleman sayısı n kadar olabilir.
//aslında yolladığımız değerimiz array bazlıdır
//yani GO bizim sayilar üzerinden yolladığımızı ARRAY olarak alıyor.
//variadic bizim gönderdiğimiz parametreleri dizi olarak alıyor.
