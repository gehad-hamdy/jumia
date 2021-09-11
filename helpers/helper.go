package helpers

import (
	"Flash/responses"
	"strings"
)

func GetRegularExpressions() map[string]responses.RegularExpressions {
	mappingRegexByCountryCode := make(map[string]responses.RegularExpressions)

	mappingRegexByCountryCode["273"] = responses.RegularExpressions{
		CountryName: "Cameroon",
		CountryCode: "273",
		Regex:"?[2368]\\d{7,8}$"}

	mappingRegexByCountryCode["251"] = responses.RegularExpressions{
		CountryName: "Ethiopia",
		CountryCode: "251",
		Regex:"?[1-59]\\d{8}$",
	}

	mappingRegexByCountryCode["212"] = responses.RegularExpressions{
		CountryName: "Morocco",
		CountryCode: "212",
		Regex:"?[5-9]\\d{8}$",
	}

	mappingRegexByCountryCode["258"] = responses.RegularExpressions{
		CountryName: "Mozambique",
		CountryCode: "258",
		Regex:"?[28]\\d{7,8}$",
	}

	mappingRegexByCountryCode["256"] = responses.RegularExpressions{
		CountryName: "Uganda",
		CountryCode: "256",
		Regex:"?\\d{9}$",
	}

	return mappingRegexByCountryCode
}

func IsPositiveInteger(number string) bool {
	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	return strings.IndexFunc(number, isNotDigit) == -1
}
