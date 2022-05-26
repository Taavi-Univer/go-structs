package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // we dont have to specify field name (contact)
}

func main() {
	// taavi := person{"Taavi", "Univer"}
	// fmt.Println(taavi.firstName)

	// taavi := person{lastName: "Univer", firstName: "Taavi"}
	// fmt.Println(taavi)

	// var taavi person

	// taavi.firstName = "Taavi"
	// taavi.lastName = "Univer"

	// fmt.Println(taavi)
	// fmt.Printf("%+v", taavi)

	taavi := person{
		firstName: "Taavi",
		lastName:  "Univer",
		contact: contactInfo{
			email:   "taaviuniver@gmail.com",
			zipCode: 80013,
		},
	}

	// taavi.print()

	// taaviPointer := &taavi
	// taaviPointer.updateName("Taav")

	taavi.updateName("Taav")
	taavi.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
