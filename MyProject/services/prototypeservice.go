package services

import (
	"Go_Project/models"
)
/*  Prototype Pattern Implimentation...

The Prototype Design Pattern is a creational design pattern that focuses on creating objects by copying or cloning
an existing object, known as the prototype, instead of creating them from scratch.
It provides a way to create new objects with the same structure and properties as an existing object,
which can be particularly useful when the cost of creating an object is high or complex.
*/

//Prototype interface consists clone() method
type Prototype interface {
	Clone()
}

//concrete Prototype which implements the clone() method
type MyEmpDetails struct {
	*models.EmpDetails
}


//implementing clone() method
func (u *MyEmpDetails) Clone() *MyEmpDetails {
	return &MyEmpDetails{
		EmpDetails: &models.EmpDetails{
			Id:      u.Id,
			Fname:   u.Fname,
			Lname:   u.Lname,
			Age:     u.Age,
			Phno:    u.Phno,
			Email:   u.Email,
			City:    u.City,
			State:   u.State,
			Country: u.Country,
			Zip:     u.Zip,
		},
	}
}


// clone object
var PrototypeEmp = &MyEmpDetails{
	EmpDetails: &models.EmpDetails{
		Id:      1,
		Fname:   "First name",
		Lname:   "Last name",
		Age:     24,
		Phno:    "+919966996115",
		Email:   "employee@gmail.com",
		City:    "Vizag",
		State:   "AP",
		Country: "India",
		Zip:     533288,
	},
}
