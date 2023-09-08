package infrastructure

import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"Go_Project/models"
	"Go_Project/config"
)

var DB *gorm.DB

func ConnectDatabase() {
    d:=config.GetDbconfig()
	dsn := d.DbUser+":"+d.DbPass+"@tcp(127.0.0.1:3306)/"+d.DbName+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db

	DB.AutoMigrate(&models.EmpDetails{})

}
