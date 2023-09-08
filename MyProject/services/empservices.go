package services

import(
	"Go_Project/models"
	"github.com/gin-gonic/gin"
	"Go_Project/repositories"
)

//To retrieve details of employees
func Details()(e []models.EmpDetails,err error){
	return repositories.Detailsrepo()
}

//to get details by id
func ById(c *gin.Context)(e models.EmpDetails,err error){
	return repositories.ByIdRepo(c)
}

//Delete by Id
func DelbyId(c *gin.Context)(msg string){
	return repositories.DelIdRepo(c)
}

//creating record
func Add(e models.EmpDetails)(s string){
	return repositories.AddRepo(e)
}

//updating

func UpRec(c *gin.Context)(msg string){
	return repositories.UpRecRepo(c)
}