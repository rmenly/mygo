package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(n1 string) *person {
	p := person{name: n1}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 21})
	fmt.Println(person{name: "Chock"})
	fmt.Println("zhen: ", &person{name: "Deck", age: 23})
	fmt.Println(newPerson("Frank"))

	s := person{name: "Eric", age: 25}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 90
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"happy",
		true,
	}
	fmt.Println(dog)

}
