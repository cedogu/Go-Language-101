package slices

import "fmt"

func Demo2() {
	cities := []string{"İstanbul", "Ankara", "İzmir"}
	fmt.Println(cities)
	citiesCopy := make([]string, len(cities))
	fmt.Println(citiesCopy)
	copy(citiesCopy, cities)
	fmt.Println(citiesCopy)
	cities = append(cities, "Muğla") //cities slice'a yeni bir slice olup, cities'e ekliyoruz. CitiesCopy'e Mugla eklenmez çünkü tek bir slice'a ekledik
	fmt.Println(cities)
	fmt.Println(citiesCopy)

	fmt.Println(cities[1:3])
	fmt.Println(cities[1:])
	fmt.Println(cities[:2])

}
