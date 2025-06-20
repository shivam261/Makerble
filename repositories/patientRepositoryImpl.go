package repositories

import (
	"github/shivam261/ClinicManagement/initializers"
	"github/shivam261/ClinicManagement/models"
)

type patientRepository struct{}

func NewPatientRepository() PatientRepository {
	return &patientRepository{}
}

func (r *patientRepository) Create(patient *models.Patient) error {
	return initializers.DB.Create(patient).Error
}

func (r *patientRepository) FindAll() ([]models.Patient, error) {
	var patients []models.Patient
	if err := initializers.DB.Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}

func (r *patientRepository) UpdateByID(id string, data models.Patient) error {
	return initializers.DB.Model(&models.Patient{}).Where("id = ?", id).Updates(data).Error
}
