package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

type ContactInfo struct {
	email   string
	zipCode int
}

func main() {
	richard := Person{"Richard", "Munoz", ContactInfo{"jim@gmail.com", 12345}}
	fmt.Println(richard)

	richard2 := Person{firstName: "Richard", lastName: "Munoz"}
	fmt.Println(richard2)

	// assigned default zero value
	var richard3 Person

	richard3.firstName = "Richard"
	richard3.lastName = "Munoz"

	fmt.Println(richard3)

	// print all field names and values
	fmt.Printf("\n%+v", richard3)

	jim := Person{
		firstName: "Jim", lastName: "Party", contact: ContactInfo{
			email: "jim@gmail.com", zipCode: 12345,
		},
	}

	jimPointer := &jim // gets memory address (reference) and assigns to jimPointer
	jimPointer.updateName("Jimmy")

	updateName2(&jim, "Jimmy")

	jim.print()

	// Testing mutability with new object assignment
	ariel := Person{firstName: "Ariel", lastName: "Munoz"}
	ariel.print()
	zorro1 := updateName3(ariel, "zorro1")
	zorro1.print()
}

func (pointerToPerson *Person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func updateName2(p *Person, newName string) {
	p.firstName = newName
}

func updateName3(p Person, newName string) Person {
	p.firstName = newName
	return p
}

func (p Person) print() {
	fmt.Printf("\n%+v", p)
}
