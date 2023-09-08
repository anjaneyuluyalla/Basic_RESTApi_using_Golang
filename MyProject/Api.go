package main

import (
	"Go_Project/infrastructure"
	"Go_Project/routers"
)

func main() {

	infrastructure.ConnectDatabase()
	
	routers.Route()
}
