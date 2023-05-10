package domain

type Call struct {
	Id          int64  `json:"id"`
	CustomerId  int64  `json:"customerId"`
	//FullName    string `json:"fullname"`
	DateOfCall  string `json:"dateOfCall"`
	TimeOfCall  string `json:"timeOfCall"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

