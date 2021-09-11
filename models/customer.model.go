package models

type Customer struct {
	Phone string  `json:"phone"`
	CountryCode string  `json:"country_code"`
}

func (customer *Customer) TableName() string {
	return "customer"
}
