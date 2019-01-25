package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	g := person{
		firstName: "gaurav",
		lastName:  "Bagul",
		contactInfo: contactInfo{
			email:   "g@gmail.com",
			zipCode: 123,
		},
	}
	// gauravPointer := &gaurav
	// fmt.println("gauravPointer-> %v",*gauravPointer)
	g.updateName("Gaurav")
	g.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
