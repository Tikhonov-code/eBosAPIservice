package domain

type Customer struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Address     string `json:"address"`
	PostCode    string `json:"postcode"`
	Country     string `json:"country"`
	DateOfBirth string `json:"dateofbirth"`
}

type CustomerLookup struct {
	Id       int64  `json:"id"`
	FullName string `json:"fullname"`
}
