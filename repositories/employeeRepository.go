package repositories

import (
	"github/shivam261/ClinicManagement/models"
)

type EmployeeRepository interface {
	Create(employee *models.Employee) error
	FindByEmail(email string) (*models.Employee, error)
}
