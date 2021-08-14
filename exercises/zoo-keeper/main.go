// Workers work inside zoo premises
// Animal keeper takes care of animals and feed them
// Stuffs will work with Visitors, they are not allowed to feed animal
package main

import "fmt"

type WorkerInterface interface {
	getName() string
}
type AnimalKeeperInterface interface {
	feedAnimal() bool
}

type OtherStuffInterface interface {
	bookTicket() bool
}

type AnimalKeeper struct {
	firstname       string
	lastname        string
	visitAnimalPass bool
}

type OtherStuff struct {
	firstname string
	lastname  string
}

func main() {
	animalKeeper := AnimalKeeper{
		firstname:       "Jim",
		lastname:        "Cobert",
		visitAnimalPass: true,
	}
	otherStuff := OtherStuff{
		firstname: "Mary",
		lastname:  "Smith",
	}

	fmt.Println("name of animalKeeper: ", animalKeeper.getName())
	fmt.Println("name of otherStuff: ", otherStuff.getName())
	fmt.Println("animalkeeper can feed animals: ", animalKeeper.feedAnimal())
	fmt.Println("otherStuff can book tickets: ", otherStuff.bookTicket())
}

// type AnimalKeeper implements WorkerInterface
func (a AnimalKeeper) getName() string {
	return a.firstname + " " + a.lastname
}

// type AnimalKeeper implements AnimalKeeperInterface
func (a AnimalKeeper) feedAnimal() bool {
	return true
}

// type OtherStuff implements WorkerInterface
func (o OtherStuff) getName() string {
	return o.firstname + " " + o.lastname
}

// type OtherStuff implements OtherStuffInterface
func (o OtherStuff) bookTicket() bool {
	return true
}
