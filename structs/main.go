package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	richard := Person{"Richard", "Munoz"}
	fmt.Println(richard)

	richard2 := Person{firstName: "Richard", lastName: "Munoz"}
	fmt.Println(richard2)

	// assigned default zero value
	var richard3 Person

	richard3.firstName = "Richard"
	richard3.lastName = "Munoz"

	fmt.Println(richard3)

	// print all field names and values
	fmt.Printf("%+v", richard3)

}
