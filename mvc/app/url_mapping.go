package app

import (
	"main/mvc/controllers"

)

func mapUrls() {
	// login section
	router.POST("/login", controllers.PostLogin)
	//router.GET("/login/:username/:password", controllers.GetLogin)

	// Calls section
	router.GET("/calls", controllers.GetCalls)
	router.DELETE("/deletecall/:id", controllers.DeleteCall)
	router.PUT("/updatecall/:id", controllers.UpdateCall)
	router.POST("/insertcall", controllers.InsertCall)

	// customers section
	router.GET("/customers", controllers.GetCustomers)
	router.DELETE("/deletecustomer/:id", controllers.DeleteCustomer)
	router.POST("/insertcustomer", controllers.InsertCustomer)
	router.PUT("/updatecustomer/:id", controllers.UpdateCustomer)
	router.GET("/customersList", controllers.GetCustomersList)

	// reports section
	router.GET("/report", controllers.GetReport)

	

}