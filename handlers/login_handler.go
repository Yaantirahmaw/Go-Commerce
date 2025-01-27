package handlers

import (
	"go-commerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		var user models.User
		if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
			c.JSON(401, gin.H{"message": "Invalid credentials"})
			return
		}

		token, err := CreateToken(user.ID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}

		c.JSON(200, gin.H{"token": token})
	}
}
