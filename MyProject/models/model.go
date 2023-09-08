package models

type EmpDetails struct {
	//	gorm.Model
	Id      int    `json:"id" gorm:"primaryKey"`
	Fname   string `json:"fname" binding:"required"`
	Lname   string `json:"lname" binding:"required"`
	Age     int    `json:"age" binding:"gte=20,lte=60"`
	Phno    string `json:"phno" binding:"required,e164"`
	Email   string `json:"email" binding:"required,email"`
	City    string `json:"city" binding:"required"`
	State   string `json:"state" binding:"required"`
	Country string `json:"country" binding:"required"`
	Zip     int    `json:"zip" binding:"gte=100000,lte=899999"`
}

type Prototype interface{
	Clone()
}

func (u *EmpDetails)Clone() *EmpDetails{
	return &EmpDetails{
		Id : u.Id,
		Fname: u.Fname,
		Lname: u.Lname,
		Age: u.Age,
		Phno: u.Phno,
		Email: u.Email,
		City: u.City,
		State: u.State,
		Country: u.Country,
		Zip: u.Zip,
	}
}