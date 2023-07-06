package main

import "fmt"

type contactInfo struct {
	email string
	phone int
}


type person struct {
	firstName string
	lastName string
	contact contactInfo
}


func  main()  {
arslan := person {
	firstName: "muhammad", lastName: "arslan", contact: contactInfo {
		email:"muhammadarslan6492@gmail.com",
		phone:6262626262622,
	},
}


arslan.updateFistName("Ali")

arslan.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateFistName(fName string){
	(*pointerToPerson).firstName = fName 
}