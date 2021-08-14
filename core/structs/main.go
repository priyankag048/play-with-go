package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex.firstName)
	// fmt.Println(alex.lastName)

	// var alex1 person
	// alex1.firstName = "Alex"
	// alex1.lastName = "Anderson"
	// fmt.Println(alex1)
	// fmt.Printf("%+v", alex1)

	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 560048,
		},
	}
	(&jim).updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}
