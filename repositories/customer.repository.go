package repositories

import (
	"Flash/models"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	connection *gorm.DB
}

func GetCustomerRepository(connection *gorm.DB) *CustomerRepository {
	return &CustomerRepository{connection : connection}
}

func (repo CustomerRepository) ListPhoneNumbers() ([]models.Customer, error) {
	var model []models.Customer
	err := repo.connection.Model(models.Customer{}).
		Select("trim(substr(phone, 1, instr(phone, ' ') - 1) , '(,)') AS country_code, substr(phone, instr(phone, ' ') + 1) AS phone").
		Scan(&model).
		Order("country_code DESC").
		Error

	return model, err
}