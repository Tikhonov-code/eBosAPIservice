package main

import (
	"main/mvc/app"
)

var DatabaseConnection string = "mongodb://localhost:27017"

func main()  {
	app.StartApp()
}




