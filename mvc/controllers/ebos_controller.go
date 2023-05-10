package controllers

import (
	"io/ioutil"
	"log"
	//"main/mvc/domain"
	"main/mvc/services"
	"main/mvc/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostLogin(c *gin.Context) {
	data :=ExtractBody(c)

	user, apiErr := services.EBOSservice.PostLogin(data)
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}


//GetCalls
func GetCalls(c *gin.Context) {
	log.Println("GetCalls")
	calls, apiErr := services.EBOSservice.GetCalls()
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, calls)
}
//DeleteCall
func DeleteCall(c *gin.Context) {
	id := c.Param("id")
	log.Printf("DeleteCall id = %s",id)
	apiErr := services.EBOSservice.DeleteCall(id)
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "DeleteCall")
}

//InsertCall
func InsertCall(c *gin.Context) {
	data :=ExtractBody(c)
	log.Println("InsertCall data = ",data)
	
	msg,apiErr := services.EBOSservice.InsertCall(data)
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, msg)
}

func UpdateCall(c *gin.Context) {
	//utils.Respond(c, http.StatusOK, "Call Updated")
	id := c.Param("id")
	requestBody, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        c.Writer.WriteHeader(http.StatusBadRequest)
        c.Writer.Write([]byte("Failed to read request body"))
        return
    }

    // Convert the byte slice to a string
    requestBodyString := string(requestBody)
	
	log.Println("requestBodyString = ",requestBodyString)
	result := Extract(requestBodyString)
	log.Println("result = ",result)

	msg,apiErr := services.EBOSservice.UpdateCall(result,id)
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, msg)
}

// GetCustomers
func GetCustomers(c *gin.Context) {
	customers, apiErr := services.EBOSservice.GetCustomers()
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, customers)
}

// GetCustomersList
func GetCustomersList(c *gin.Context) {
	customers, apiErr := services.EBOSservice.GetCustomersList()
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, customers)
}

//DeleteCustomer
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	log.Printf("DeleteCustomer id = %s",id)

	msg,apiErr := services.EBOSservice.DeleteCustomer(id)
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, msg)
}
//InsertCustomer
func InsertCustomer(c *gin.Context) {
	data :=ExtractBody(c)
	log.Println("InsertCustomer data = ",data)
	
	msg,apiErr := services.EBOSservice.InsertCustomer(data)
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, msg)
}
//UpdateCustomer
func UpdateCustomer(c *gin.Context) {
	//utils.Respond(c, http.StatusOK, "Call Updated")
	id := c.Param("id")
	requestBody, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        c.Writer.WriteHeader(http.StatusBadRequest)
        c.Writer.Write([]byte("Failed to read request body"))
        return
    }

    // Convert the byte slice to a string
    requestBodyString := string(requestBody)
	
	log.Println("requestBodyString = ",requestBodyString)
	result := Extract(requestBodyString)
	log.Println("result = ",result)

	msg,apiErr := services.EBOSservice.UpdateCustomer(result,id)
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, msg)
}



// GetReports
func GetReport(c *gin.Context) {
	reports, apiErr := services.EBOSservice.GetReport()
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, reports)
}

func ExtractBody (c *gin.Context) string {
	requestBody, _ := ioutil.ReadAll(c.Request.Body)
    requestBodyString := string(requestBody)
	return requestBodyString
}

func Extract (body string) string {
	var result string
	var start bool
	for _, c := range body {
		if c == '{' {
			start = true
		}
		if start {
			result += string(c)
		}
	}
	return result
}