package controllers

import (
	"github/shivam261/ClinicManagement/initializers"
	"os"
	"time"

	"github/shivam261/ClinicManagement/models"
	"github/shivam261/ClinicManagement/repositories"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// this function wiil register doctor and receptionists
	var body struct {
		Role            string `json:"role"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConformPassword string `json:"conformPassword"`
		Name            string `json:"name"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if body.Role != "doctor" && body.Role != "receptionist" {
		c.JSON(400, gin.H{"error": body.Role + " is not a valid role"})
		return
	}
	if body.Password != body.ConformPassword {
		c.JSON(400, gin.H{"error": "Passwords do not match"})
		return
	}
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}
	body.Password = string(hashedPassword)

	// Create a new user in the database
	employee := &models.Employee{
		Role:     body.Role,
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
	}
	if err := repositories.NewEmployeeRepository().Create(employee); err != nil {
		c.JSON(409, gin.H{"error": "Failed to register " + body.Role + ". Try using a different email"})
		return
	}
	c.JSON(201, gin.H{"message": body.Role + " registered successfully"})

}
func Login(c *gin.Context) {

	// this function will login doctor and receptionists
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var employee models.Employee

	if emp, err := repositories.NewEmployeeRepository().FindByEmail(body.Email); err != nil {

		c.JSON(404, gin.H{"error": "User not found"})
		return
	} else {
		employee = *emp
	}

	if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(body.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": employee.Role,
		"exp":  time.Now().Add(time.Minute * 1).Unix(), // Token expires in 1 minutes
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}
	initializers.RedisClient.Set(initializers.Ctx, tokenString, employee.Role, time.Minute*1)
	println("Token stored in Redis with key:", initializers.RedisClient.Get(initializers.Ctx, tokenString).Val())
	c.SetCookie("Authorization", tokenString, 3600, "/", "localhost", false, true)

	c.JSON(200, gin.H{"message": "Login successful", "role": employee.Role})

}
func Logout(c *gin.Context) {
	// this function will logout the user by clearing the cookie
	c.SetCookie("Authorization", "", -1, "/", "localhost", true, true)
	c.JSON(200, gin.H{"message": "Logout successful"})
}
