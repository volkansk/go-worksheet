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
	//var volkan person
	// volkan.firstName = "Volkan"
	// volkan.lastName = "Isik"
	// volkan.contact = contactInfo{
	// 	email:   "volkanserifisik@gmail.com",
	// 	zipCode: 34000,
	// }

	volkan := person{
		firstName: "Volkan",
		lastName:  "Isik",
		contact: contactInfo{
			email:   "volkanserifisik@gmail.com",
			zipCode: 34000,
		},
	}

	//volkan.updateName("Volkan Serif") // volkan is not updated because struct received in receiver method is copied to different address in memory. must use pointer

	// volkanPointer := &volkan // &variable -> give me the memory address of the value this variable is pointing at
	// volkanPointer.updateName("Volkan Serif")

	volkan.updateName("Volkan Serif") //shorcut pointer usage
	volkan.print()
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (volkanPointer *person) updateName(newFirstName string) {
	(*volkanPointer).firstName = newFirstName //*pointer -> give me the value this memory address is pointing at
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
