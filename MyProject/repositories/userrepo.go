package repositories

import (
	"Go_Project/infrastructure"
	"Go_Project/models"
	"github.com/gin-gonic/gin"
	"strconv"

)

//details..
func Detailsrepo()(emp []models.EmpDetails,err error){
	var employee []models.EmpDetails
	err = infrastructure.DB.Find(&employee).Error
	return employee,err
}

//details by id

func ByIdRepo(c *gin.Context)(emp models.EmpDetails,err error){
	empId, _ := strconv.Atoi(c.Param("id"))
	err = infrastructure.DB.First(&emp, empId).Error
	return emp,err
}

//delete by id

func DelIdRepo(c *gin.Context)(msg string){
	var employee models.EmpDetails
	empId,_:=strconv.Atoi(c.Param("id"))
	if err := infrastructure.DB.First(&employee, empId).Error; err != nil {
		return "Not ok"
	}
	infrastructure.DB.Delete(&employee)
	return "Ok"
}

// adding details(post operation)

func AddRepo(e models.EmpDetails)(msg string){
	if err := infrastructure.DB.Create(&e).Error; err != nil {
		return "Not Ok"
	}
	return "Ok"
}

//updating..

func UpRecRepo(c *gin.Context)(msg string){
	var employee models.EmpDetails
	empId, _ := strconv.Atoi(c.Param("id"))
	if err := infrastructure.DB.First(&employee, empId).Error; err != nil {
		return "Error1"
	}
	var updetail models.EmpDetails
	if err := c.ShouldBindJSON(&updetail); err != nil {
		return "Error2"
	}
	employee=updetail
	infrastructure.DB.Save(&employee)
	return "Ok"
}