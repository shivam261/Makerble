package unit

import (
	"bytes"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github/shivam261/ClinicManagement/controllers"
)

func setupPatientRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/doctor/patients", controllers.GetAllPatients)
	r.POST("/receptionist/registerPatient", controllers.AddPatient)
	r.PUT("/doctor/patients/:id", controllers.UpdatePatientById)
	return r
}

func TestGetAllPatients_Empty(t *testing.T) {
	r := setupPatientRouter()
	req, _ := http.NewRequest("GET", "/doctor/patients", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "[")
}

func TestAddPatient_InvalidInput(t *testing.T) {
	r := setupPatientRouter()
	invalidBody := `{"age": 45}` // Missing required fields

	req, _ := http.NewRequest("POST", "/receptionist/registerPatient", bytes.NewBufferString(invalidBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Name and Disease fields are required")
}

func TestUpdatePatientById_InvalidInput(t *testing.T) {
	r := setupPatientRouter()
	invalidUpdate := `{"age": "invalid"}`

	req, _ := http.NewRequest("PUT", "/doctor/patients/1", bytes.NewBufferString(invalidUpdate))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid input")
}
