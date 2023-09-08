package controllers

import (
	"Go_Project/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Welcome page
func Open(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// Get all details of employees
func GetDetails(c *gin.Context) {
	var Data, err = services.Details()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Details not found"})
		return
	}
	c.JSON(200, Data)
}

// Get details of specific employee
func GetById(c *gin.Context) {
	var Data, err = services.ById(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Details not found"})
		return
	}
	c.JSON(http.StatusOK, Data)
}

// Delete details of an employee
func DelRecord(c *gin.Context) {
	var m = services.DelbyId(c)
	if m != "Ok" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, "User deleted successfully")
}

// Add details of an employee
func AddRecord(c *gin.Context) {
	//here creating employee object using clone method(Prototype pattern)
	
	employee := services.PrototypeEmp.EmpDetails.Clone()

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Details not sufficient..."})
		return
	}
	var m = services.Add(*employee)
	if m != "Ok" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable create"})
		return
	}
	c.JSON(http.StatusCreated, employee)

}

// Update employee details
func UpdateRecord(c *gin.Context) {
	var m = services.UpRec(c)
	if m == "Ok" {
		c.JSON(http.StatusOK, gin.H{"message": "Details updated..."})
	} else if m == "Error1" {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found..."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Insufficient data..."})
	}
}
