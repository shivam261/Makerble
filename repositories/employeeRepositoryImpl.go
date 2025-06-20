package repositories

import (
	"github/shivam261/ClinicManagement/initializers"
	"github/shivam261/ClinicManagement/models"
)

type employeeRepository struct{}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{}
}

func (r *employeeRepository) Create(employee *models.Employee) error {
	return initializers.DB.Create(employee).Error
}

func (r *employeeRepository) FindByEmail(email string) (*models.Employee, error) {
	var emp models.Employee
	if err := initializers.DB.Where("email = ?", email).First(&emp).Error; err != nil {
		return nil, err
	}
	return &emp, nil
}
