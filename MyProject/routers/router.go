package routers

import(
	"github.com/gin-gonic/gin"
	"Go_Project/controllers"
	"Go_Project/middleware"

)

func Route(){
	app := gin.Default()
	app.LoadHTMLGlob("views/*.html")
	app.GET("/",middleware.Authenticate,controllers.Open)
	app.GET("/employees", controllers.GetDetails)
	app.GET("/employees/:id",controllers.GetById)
	app.DELETE("/employees/:id",controllers.DelRecord)
	app.POST("/employees",controllers.AddRecord)
	app.PUT("/employees/:id", controllers.UpdateRecord)
	app.Run()
}