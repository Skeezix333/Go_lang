package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func main() {
	//alex := person {"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// contact: contactInfo{
		// 	email:   "jim@gmail.com",
		// 	zipCode: 94000,
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// fmt.Printf("%+v", jim)

	//& turns jim into a pointer to its memory location
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	// jim.print()
	jim.updateName("jimmy")
	jim.print()
}

//rewatch lec. 44
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
