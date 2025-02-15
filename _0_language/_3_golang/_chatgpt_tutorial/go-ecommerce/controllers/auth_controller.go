package controllers

import (
	"fmt"
	"go-ecommerce/db"
	"go-ecommerce/models"
	"go-ecommerce/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	fmt.Println("📌 Hashed Password (Before Saving to DB):", string(hashedPassword))

	user.Password = string(hashedPassword)

	// Debug: Print the hashed password
	fmt.Println("Hashed Password:", user.Password)

	// Save user in DB
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.DB.Where("email = ?", requestBody.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	// fmt.Println("📌 Testing Hash Manually:")

	// testPassword := "12345" // The password you entered during signup
	// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(testPassword))

	// if err == nil {
	// 	fmt.Println("✅ TEST PASSED: Password '12345' matches the stored hash.")
	// } else {
	// 	fmt.Println("🚨 TEST FAILED: Password '12345' does NOT match the stored hash.", err)
	// }

	// Debug: Print stored password vs entered password
	fmt.Println("📌 Stored Hashed Password from DB:", user.Password)
	fmt.Println("📌 Entered Password:", requestBody.Password)

	// Check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))
	fmt.Println("Comparison Error:", err) // Print comparison error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
