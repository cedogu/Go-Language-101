package interfaces

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12

}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0
	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "İstanbul / Kadıköy"}
	credit2 := Mortgage{rate: 12, creditPaymentTotal: 50000, address: "Muğla / Bodrum"}
	credit3 := Car{rate: 15, creditPaymentTotal: 60000, carInfo: "Polo/Germany"}
	credit4 := Car{rate: 20, creditPaymentTotal: 50000, carInfo: "Wolkswagen / Germany"}

	credits := []CreditCalculator{credit1, credit2, credit3, credit4}
	total := CalculateMonthlyPayment(credits)
	fmt.Println("Total payment: ", total)
}
