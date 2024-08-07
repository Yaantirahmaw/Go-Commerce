package handlers

import (
	"go-commerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		var existingUser models.User
		if err := db.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
			c.JSON(400, gin.H{"message": "Username already exists"})
			return
		}

		newUser := models.User{
			Username: input.Username,
			Password: input.Password, 
		}

		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}

		token, err := CreateToken(newUser.ID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}

		c.JSON(200, gin.H{"message": "User registered successfully", "token": token})
	}
}
