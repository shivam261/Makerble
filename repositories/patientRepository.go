package repositories

import (
	"github/shivam261/ClinicManagement/models"
)

type PatientRepository interface {
	Create(patient *models.Patient) error
	FindAll() ([]models.Patient, error)
	UpdateByID(id string, data models.Patient) error
}
