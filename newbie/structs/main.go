package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	hasan := person{
		firstName: "hasan",
		lastName: "kahraman",
		contactInfo: contactInfo{
			email: "kahraman.h@outlook.com",
			zipCode: 34000,
		},
	}

	// hasanPointer := &hasan
	// hasanPointer.updateEmail("hassohassan@hotmail.com")
	//or
	hasan.updateEmail("hassohassan@hotmail.com")
	//that because your updateEmail parameter's got "*"  --> (*person)



	// hasan.print()

	createSlice()
}

func (pointerToPerson *person) updateEmail(newEmail string) {
	(*pointerToPerson).contactInfo.email = newEmail
}

func (p person) print() {
	fmt.Printf("%+v", p)	
}

//////// practice slices

func createSlice() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	// Whenever you pass an integer, float, string, or struct into a function, go creates a copy of each argument, and these copies used inside of function.
	// But slice does not work like that. 
	// There is still a copy but both copies point same memory, same array. 
	fmt.Println(mySlice)

}

func updateSlice(s []string) {
	s[0] = "Bye"
}

