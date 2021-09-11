package services

import (
	"Flash/helpers"
	"Flash/logger"
	"Flash/models"
	"Flash/pkg/db"
	"Flash/repositories"
	"Flash/responses"
	"strconv"
	"strings"
)

type CustomerService struct {
	 customerRepository *repositories.CustomerRepository
}

func GetCustomerService() CustomerService {
	conn, _ := db.GetDB()
	return CustomerService{
		customerRepository: repositories.GetCustomerRepository(conn),
	}
}

func (service CustomerService) ListPhoneNumbers() []models.Customer {
	model, err := service.customerRepository.ListPhoneNumbers()
	if err != nil {
		logger.Info(err.Error())
		println(err.Error())
	}
	return model
}

func (service CustomerService) CheckValidationNumber(phoneNumbers []models.Customer) []responses.PhoneNumberStatusResult {
    patternExpression := helpers.GetRegularExpressions()
	var data []responses.PhoneNumberStatusResult
	
	for _, phoneNumber := range phoneNumbers {
		if val, ok := patternExpression[phoneNumber.CountryCode]; ok {
			 result := responses.PhoneNumberStatusResult{}
			 result.PhoneNumber = phoneNumber.Phone
			 result.Country = val.CountryName
			 result.CountryCode = val.CountryCode
			if helpers.IsPositiveInteger(phoneNumber.Phone) {
				if validSquareBracketsNumber(val.Regex, phoneNumber.Phone) {
					if validConsecutiveCharacters(val.Regex, phoneNumber.Phone) {
						result.State = "OK"
					} else {
						result.State = "NO"
					}
				} else {
					result.State = "NO"
				}
			} else {
				result.State = "NO"
			}
			data = append(data, result)
		}
	}
	return data
}

func validSquareBracketsNumber(pattern string, number string) bool {
	indexOfOpenSquareBrackets := strings.Index(pattern, "[")
	indexOfCloseSquareBrackets := strings.Index(pattern, "]")
	if indexOfOpenSquareBrackets > 0 && indexOfCloseSquareBrackets > 0 {
		squareBrackets := pattern[indexOfOpenSquareBrackets + 1 : indexOfCloseSquareBrackets]
		if helpers.IsPositiveInteger(squareBrackets) {
            return checkIncludedElements(number, squareBrackets)
		} else {
			validRange := checkValidRange(squareBrackets, number)
			return validRange
		}
	}
	return true
}

func validConsecutiveCharacters(pattern string, number string) bool {
	indexOfOpenBrackets := strings.Index(pattern, "{")
	indexOfCloseBrackets := strings.Index(pattern, "}")
	if indexOfOpenBrackets > 0 && indexOfCloseBrackets > 0 {
		lengthValue := pattern[indexOfOpenBrackets + 1 : indexOfCloseBrackets]
		//only one number
		if len(lengthValue) == 1 {
			lengthValueInt, _ := strconv.Atoi(lengthValue)
			if len(number) == lengthValueInt {
				return true
			}else {
				return false
			}
		} else {
			//multiple range
			validNumber := checkValidRange(lengthValue, number)

			return validNumber
		}
	}
	return true
}

func checkValidRange(squareBrackets string, number string) bool {
	rangeNumber1, _ := strconv.Atoi(squareBrackets[0 : 1])
	rangeNumber2, _ := strconv.Atoi(squareBrackets[2 : 3])
	validRange := false
	for _, char := range number {
		intVar, _ := strconv.Atoi(string(char))
		if intVar >= rangeNumber1 && intVar <= rangeNumber2 {
			validRange = true
		} else {
			validRange = false
			break
		}
	}

	return validRange
}

func checkIncludedElements(number, regex string) bool {
	validNum := true
	for _, x := range number {
		_, found := Find(regex, string(x))
		if !found {
			validNum = false
			break
		} else {
			validNum = true
		}
	}
	return validNum
}

func Find(regex string, num string) (int, bool) {
	for i, item := range regex {
		if string(item) == num {
			return i, true
		}
	}
	return -1, false
}