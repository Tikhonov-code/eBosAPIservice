package services

import (
	"fmt"
	"strings"
	//"log"
	//"main/mvc/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostLogin(t *testing.T) {
	ebss := eBOSservice{}
	user, err := ebss.PostLogin("username=manager&password=e10adc3949ba59abbe56e057f20f883e")
	assert.Nil(t, err)
	assert.NotNil(t, user)

	fmt.Printf("Name =%s  Role= %s  Status=%s", user.Name, user.Role, user.Status)
}

func TestPostLoginIllegal(t *testing.T) {
	ebss := eBOSservice{}
	user, err := ebss.PostLogin("{'name': 'b@a.a', 'password': 'e10adc3949ba59abbe56e057f20f883e'}")
	assert.NotNil(t, err)
	assert.EqualValues(t, user.Status, "illegal")

	fmt.Printf("Code =%s  Message= %s  Status=%d", err.Code, err.Message, err.StatusCode)
}

func TestGetLoginEmployeeLegal(t *testing.T) {
	ebss := eBOSservice{}
	user, err := ebss.PostLogin("{'name': 'a@a.a', 'password': 'e10adc3949ba59abbe56e057f20f883e'}")
	assert.Nil(t, err)
	assert.EqualValues(t, user.Status, "legal")

}

// OpenConnection
func TestOpenConnection(t *testing.T) {
	sqlcon, err := OpenConnection()
	assert.Nil(t, err)
	assert.EqualValues(t, sqlcon.Stats().OpenConnections, 0)

}

// GetCalls
func TestGetCalls(t *testing.T) {
	ebss := eBOSservice{}
	calls, err := ebss.GetCalls()
	assert.Nil(t, err)
	assert.NotNil(t, calls)

	for _, v := range calls {
		fmt.Printf("ID =%d  Name= %s  DateOfCall=%s  TimeOfCall=%s  Subject=%s Description=%s", v.Id, v.FullName, v.DateOfCall, v.TimeOfCall, v.Subject, v.Description)
	}

}

// DeleteCall
func TestDeleteCall(t *testing.T) {
	ebss := eBOSservice{}
	err := ebss.DeleteCall("1")
	assert.Nil(t, err)

}

// UpdateCall
func TestUpdateCall(t *testing.T) {
	ebss := eBOSservice{}
	updateCall := "{\"description\":\"success test updated value6666666\"}"

	msg, err := ebss.UpdateCall(updateCall, "2")
	assert.Nil(t, err)
	assert.EqualValues(t, msg, "Call updated successfully")

}

// InsertCall
func TestInsertCall(t *testing.T) {
	ebss := eBOSservice{}
	insertCall := "{\"dateOfCall\":\"2023-05-05\",\"timeOfCall\":\"11:19\",\"subject\":\"222 test new value inserted\",\"description\":\"222 test inserted value\"}"
	msg,err := ebss.InsertCall(insertCall)
	assert.Nil(t, err)
	assert.EqualValues(t, msg, "inserted")

}

// customers section
func TestGetCustomers(t *testing.T) {
	ebss := eBOSservice{}
	customers, err := ebss.GetCustomers()
	assert.Nil(t, err)
	assert.NotNil(t, customers)

	for _, v := range customers {
		fmt.Printf("id = %d, name= %s, surname= %s, address= %s, postCode= %s, country= %s, dateOfBirth= %s\n", v.Id, v.Name, v.Surname, v.Address, v.PostCode, v.Country, v.DateOfBirth)
	}

}

// reports section
func TestGetReport(t *testing.T) {
	ebss := eBOSservice{}
	reports, err := ebss.GetReport()
	assert.Nil(t, err)
	assert.NotNil(t, reports)

	for _, v := range reports {
		fmt.Printf("Customer= %s, DateOfCall =%s, TimeOfCall=%s, Subject=%s, Description=%s\n", v.Customer, v.DateOfCall, v.TimeOfCall, v.Subject, v.Description)

	}

}

// insert customer
func TestInsertCustomer(t *testing.T) {
	ebss := eBOSservice{}
	data := "{\"name\":\"test\",}"

	msg, err := ebss.InsertCustomer(data)
	assert.Nil(t, err)
	assert.EqualValues(t, msg, "inserted")

}

func TestExtractFieldsAndValues(t *testing.T) {
	data :=`{"name":"Boban","surname":"Stojanovski","address":"address","postCode":"1000","country":"Macedonia","dateOfBirth":"2020-01-01"}`

	fields, values := ExtractFieldsAndValues(data)
	fields  = strings.ReplaceAll(fields,"\"","")
	assert.NotNil(t, fields)
	assert.NotNil(t, values)

	fmt.Printf("fields = %s\n", fields)
	fmt.Printf("values = %s\n", values)

}
