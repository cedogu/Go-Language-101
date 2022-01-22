package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("Customer Added:", c.firstName, c.lastName, c.age)
}

func (c customer) update() {
	fmt.Println("Customer updated: ", c.firstName, c.lastName, c.age)
}
func Demo2() {
	c := customer{firstName: "Doruk", lastName: "Gürsoy", age: 26}
	c.save()
	c.update()

}
