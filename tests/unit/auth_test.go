package unit

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github/shivam261/ClinicManagement/controllers"
	"github/shivam261/ClinicManagement/tests/unit/testutils"
)

func TestMain(m *testing.M) {
	testutils.SetupTestDB()
	m.Run()
}
func setupAuthRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/logout", controllers.Logout)
	return r
}

func TestRegister_Success(t *testing.T) {
	r := setupAuthRouter()
	body := map[string]string{
		"role":            "doctor",
		"email":           "doctor@example.com",
		"password":        "pass1234",
		"conformPassword": "pass1234",
		"name":            "Dr. Smith",
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Doctor registered successfully")
}

func TestRegister_PasswordMismatch(t *testing.T) {
	r := setupAuthRouter()
	body := map[string]string{
		"role":            "receptionist",
		"email":           "reception@example.com",
		"password":        "pass1234",
		"conformPassword": "wrongpass",
		"name":            "Receptionist A",
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Passwords do not match")
}

func TestLogin_InvalidUser(t *testing.T) {
	r := setupAuthRouter()
	body := map[string]string{
		"email":    "nouser@example.com",
		"password": "somepass",
	}
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "User not found")
}

func TestLogout_Success(t *testing.T) {
	r := setupAuthRouter()
	req, _ := http.NewRequest("POST", "/auth/logout", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Logout successful")
	assert.Contains(t, w.Header().Get("Set-Cookie"), "Authorization=;")
}
