package domain

type Report struct {
    Id int64 `json:"id"`
	Customer string `json:"customer"`
    DateOfCall string `json:"dateOfCall"`
    TimeOfCall string `json:"timeOfCall"`
    Subject string `json:"subject"`
    Description string `json:"description"`
}