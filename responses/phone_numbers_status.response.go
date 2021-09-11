package responses

type PhoneNumberStatusResult struct {
	Country     string `json:"country"`
	State       string `json:"state"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}
