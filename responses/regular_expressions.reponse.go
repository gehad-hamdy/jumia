package responses

type RegularExpressions struct {
	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`
	Regex       string `json:"regex"`
}
