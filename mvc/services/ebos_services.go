package services

import (
	"database/sql"
	//"encoding/json"
	"log"
	"main/mvc/domain"
	"main/mvc/utils"
	"net/http"
	"net/url"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
)

type eBOSservice struct{}

var (
	EBOSservice             eBOSservice
	StagingConnectionString = "Server=62.109.17.140;Database=EbosDB;User Id=tix;Password=1325Nikita;"
)

func OpenConnection() (*sql.DB, *utils.ApplicationError) {

	conn, err := sql.Open("sqlserver", StagingConnectionString)
	if err != nil {
		log.Println(err.Error(), sql.ErrConnDone.Error())
		return nil, &utils.ApplicationError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "Rows_error",
		}
	}

	return conn, nil
}

func (*eBOSservice) PostLogin(data string) (domain.User, *utils.ApplicationError) {

	userDb := domain.User{}
	log.Println("PostLogin data = ", data)

	values, err := url.ParseQuery(data)
	if err != nil {
		return userDb, &utils.ApplicationError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "Rows_error",
		}
	}

	username := values.Get("username")
	password := values.Get("password")

	log.Println("username = ", username, " password = ", password)

	conn, err1 := OpenConnection()
	if err1 != nil {
		return userDb, err1
	}
	defer conn.Close()

	//exec SP_GetUserByName @name='a@a.a', @password='e10adc3949ba59abbe56e057f20f883e'
	query := "exec SP_GetUserByName @name='" + username + "', @password='" + password + "'"
	log.Println("query = ", query)

	rows, err2 := conn.Query(query)

	if err2 != nil {
		return userDb, &utils.ApplicationError{
			Message:    err2.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       sql.ErrNoRows.Error(),
		}
	}
	defer rows.Close()

	var (
		id         int64
		name       string
		passwordDb string
		role       string
	)

	for rows.Next() {

		err := rows.Scan(&id, &name, &passwordDb, &role)

		if err != nil {
			log.Println(err.Error())
			return userDb, &utils.ApplicationError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
				Code:       "Row_Reading_Error",
			}
		}
		//log.Println(stage, number)
		userDb = domain.User{Id: id, Name: name, Password: passwordDb, Role: role}
	}

	if username == userDb.Name && password == userDb.Password {
		userDb.Status = "legal"

		log.Println("userDb = ", userDb)

		return userDb, nil
	} else {
		return userDb, &utils.ApplicationError{
			Message:    "User not found",
			StatusCode: http.StatusNotFound,
			Code:       "User_Not_Found",
		}
	}
}

// calls
func (*eBOSservice) GetCalls() ([]domain.Call, *utils.ApplicationError) {

	conn, err := OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := "EXEC SP_GetCallsList"

	rows, err1 := conn.Query(query)

	if err1 != nil {

		return nil, &utils.ApplicationError{
			Message:    err1.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       sql.ErrNoRows.Error(),
		}
	}
	defer rows.Close()

	var (
		id          int64
		customerId  int64
		dateOfCall  string
		timeOfCall  string
		subject     string
		description string
	)

	calls := []domain.Call{}

	for rows.Next() {

		err := rows.Scan(&id, &customerId, &dateOfCall, &timeOfCall, &subject, &description)

		if err != nil {
			log.Println(err.Error())
			return nil, &utils.ApplicationError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
				Code:       "Row_Reading_Error",
			}
		}
		//log.Println(stage, number)
		calls = append(calls, domain.Call{Id: id, CustomerId: customerId, DateOfCall: dateOfCall, TimeOfCall: timeOfCall, Subject: subject, Description: description})

	}
	err2 := rows.Err()
	if err2 != nil {
		log.Println(err2.Error())
		return nil, &utils.ApplicationError{
			Message:    err2.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "Rows_error",
		}
	}

	return calls, nil
}

// delete call
func (*eBOSservice) DeleteCall(id string) *utils.ApplicationError {

	conn, err := OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	query := "[dbo].[SP_DeleteRowById] @tablename='Calls', @id='" + id + "'"

	_, err1 := conn.Exec(query)

	if err1 != nil {

		return &utils.ApplicationError{
			Message:    err1.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       sql.ErrNoRows.Error(),
		}
	}

	return nil
}

// UpdateCall
func (*eBOSservice) UpdateCall(data string, id string) (string, *utils.ApplicationError) {
	conn, err := OpenConnection()
	if err != nil {
		return "No connection", &utils.ApplicationError{
			Message:    "No connection to database",
			StatusCode: 200,
			Code:       "200",
		}
	}
	defer conn.Close()

	query := PrepareUpdateQuery(data)
	log.Println("UpdateCall  data ="+data)
	log.Println("UpdateCall  query ="+query)

	query = "Exec SP_UpdateCallById @query='" + query + "',@id=" + id
	log.Println("UpdateCall  Query ="+query)

	_, err1 := conn.Query(query)

	if err1 != nil {

		return "", &utils.ApplicationError{
			Message:    err1.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       sql.ErrNoRows.Error(),
		}
	}
	return "call updated",nil
}

// InsertCall
func (*eBOSservice) InsertCall(data string) (string,*utils.ApplicationError) {
	fieldsStr,valuesStr := ExtractFieldsAndValues(data)
	
	fieldsStr  = strings.ReplaceAll(fieldsStr,"\"","")
	valuesStr  = strings.ReplaceAll(valuesStr,"\"","''")
	log.Println("fields ="+fieldsStr)
	log.Println("valuesStr ="+valuesStr)

	conn, err := OpenConnection()
	if err != nil {
		return "No connection",err
	}
	defer conn.Close()

	query := "SP_InsertCall @fields='"+fieldsStr+"', @values='"+valuesStr+"'" 
	query = strings.ReplaceAll(query,"fullName","customerid")

	log.Println("Query ="+query)

	_, err1 := conn.Exec(query)

	if err1 != nil {
		return "",&utils.ApplicationError{
			Message:    err1.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       sql.ErrNoRows.Error(),
		}
	}
	return "inserted",nil
}

// GetCustomers
func (*eBOSservice) GetCustomers() ([]domain.Customer, *utils.ApplicationError) {
	conn, err := OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := "Exec SP_GetCustomers"

	rows, err1 := conn.Query(query)

	if err1 != nil {

		return nil, &utils.ApplicationError{
			Message:    err1.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       sql.ErrNoRows.Error(),
		}
	}
	defer rows.Close()

	var (
		id          int64
		name        string
		surname     string
		address     string
		postCode    string
		country     string
		dateOfBirth string
	)

	customers := []domain.Customer{}

	for rows.Next() {

		err := rows.Scan(&id, &name, &surname, &address, &postCode, &country, &dateOfBirth)

		if err != nil {
			log.Println(err.Error())
			return nil, &utils.ApplicationError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
				Code:       "Row_Reading_Error",
			}
		}
		//log.Println(stage, number)
		customers = append(customers, domain.Customer{Id: id, Name: name, Surname: surname,
			Address: address, PostCode: postCode, Country: country, DateOfBirth: dateOfBirth})
	}
	err2 := rows.Err()
	if err2 != nil {
		log.Println(err2.Error())
		return nil, &utils.ApplicationError{
			Message:    err2.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "Rows_error",
		}
	}

	return customers, nil
}
//GetCustomersList
func (*eBOSservice) GetCustomersList() ([]domain.CustomerLookup, *utils.ApplicationError) {
	conn, err := OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := "Exec SP_GetCustomersList"

	rows, err1 := conn.Query(query)

	if err1 != nil {

		return nil, &utils.ApplicationError{
			Message:    err1.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       sql.ErrNoRows.Error(),
		}
	}
	defer rows.Close()

	var (
		id          int64
		fullname        string
	)

	customers := []domain.CustomerLookup{}

	for rows.Next() {

		err := rows.Scan(&id, &fullname)

		if err != nil {
			log.Println(err.Error())
			return nil, &utils.ApplicationError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
				Code:       "Row_Reading_Error",
			}
		}
		//log.Println(stage, number)
		customers = append(customers, domain.CustomerLookup{Id: id, FullName: fullname})
	}
	err2 := rows.Err()
	if err2 != nil {
		log.Println(err2.Error())
		return nil, &utils.ApplicationError{
			Message:    err2.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "Rows_error",
		}
	}

	return customers, nil

}
// DeleteCustomer
func (*eBOSservice) DeleteCustomer(id string) (string, *utils.ApplicationError) {
	conn, err := OpenConnection()
	if err != nil {
		return "No connection",err
	}
	defer conn.Close()

	query := "[dbo].[SP_DeleteCustomerById] @Id=" + id 

	_, err1 := conn.Exec(query)

	if err1 != nil {

		return "",&utils.ApplicationError{
			Message:    err1.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       sql.ErrNoRows.Error(),
		}
	}

	return "deleted",nil
}
// InsertCustomer
func (*eBOSservice) InsertCustomer(data string) (string,*utils.ApplicationError) {
	
	fieldsStr,valuesStr := ExtractFieldsAndValues(data)
	
	fieldsStr  = strings.ReplaceAll(fieldsStr,"\"","")
	valuesStr  = strings.ReplaceAll(valuesStr,"\"","''")
	log.Println("fields ="+fieldsStr)
	log.Println("valuesStr ="+valuesStr)

	conn, err := OpenConnection()
	if err != nil {
		return "No connection",err
	}
	defer conn.Close()

	query := "SP_InsertCustomer @fields='"+fieldsStr+"', @values='"+valuesStr+"'" 

	log.Println("Query ="+query)

	_, err1 := conn.Exec(query)

	if err1 != nil {
		return "",&utils.ApplicationError{
			Message:    err1.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       sql.ErrNoRows.Error(),
		}
	}
	return "inserted",nil
}
// UpdateCustomer
func (*eBOSservice) UpdateCustomer(data string, id string) (string,*utils.ApplicationError) {
	conn, err := OpenConnection()
	if err != nil {
		return "No connection", &utils.ApplicationError{
			Message:    "No connection to database",
			StatusCode: 200,
			Code:       "200",
		}
	}
	defer conn.Close()

	query := PrepareUpdateQuery(data)
	log.Println("UpdateCustomer  data ="+data)
	log.Println("UpdateCustomer  query ="+query)

	query = "Exec SP_UpdateCustomerById @query='" + query + "',@id=" + id
	log.Println("UpdateCustomer  Query ="+query)

	_, err1 := conn.Query(query)

	if err1 != nil {

		return "", &utils.ApplicationError{
			Message:    err1.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       sql.ErrNoRows.Error(),
		}
	}
	return "updated",nil
}


func  ExtractFieldsAndValues(data string) (string,string) {
	fields := ""
	values := ""
	fstr := strings.ReplaceAll(data,"{","")
	fstr  = strings.ReplaceAll(fstr,"}","")

	for _,v := range strings.Split(fstr,","){
		fields += v[:strings.Index(v,":")] + ","
		values += v[strings.Index(v,":")+1:] + ","
	}
	return fields[:len(fields)-1],values[:len(values)-1]
}

func PrepareUpdateQuery (data string) string {
	result := ""
	field := ""
	value := ""
	fstr := strings.ReplaceAll(data,"{","")
	fstr  = strings.ReplaceAll(fstr,"}","")

	for _,v := range strings.Split(fstr,","){
		field = v[:strings.Index(v,":")] 
		value = v[strings.Index(v,":")+1:]
		result += strings.ReplaceAll(field,"\"","")	+"=''"+strings.ReplaceAll(value,"\"","")+"'',"
	}


	return result[:len(result)-1]
}
	
// GetReports
func (*eBOSservice) GetReport() ([]domain.Report, *utils.ApplicationError) {
	conn, err := OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := "Exec SP_GetReport"

	rows, err1 := conn.Query(query)

	if err1 != nil {

		return nil, &utils.ApplicationError{
			Message:    err1.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       sql.ErrNoRows.Error(),
		}
	}
	defer rows.Close()

	var (
		id          int64
		customer    string
		dateOfCall  string
		timeOfCall  string
		subject     string
		description string
	)

	reports := []domain.Report{}

	for rows.Next() {

		err := rows.Scan(&id, &customer, &dateOfCall, &timeOfCall, &subject, &description)

		if err != nil {
			log.Println(err.Error())
			return nil, &utils.ApplicationError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
				Code:       "Row_Reading_Error",
			}
		}
		//log.Println(stage, number)
		reports = append(reports, domain.Report{Id: id, Customer: customer, DateOfCall: dateOfCall,
			TimeOfCall: timeOfCall, Subject: subject, Description: description})
	}
	err2 := rows.Err()
	if err2 != nil {
		log.Println(err2.Error())
		return nil, &utils.ApplicationError{
			Message:    err2.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "Rows_error",
		}
	}

	return reports, nil
}
