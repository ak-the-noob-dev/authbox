package handlers

import (
	"net/http"

	"github.com/ak-the-noob-dev/authbox/authbox/ctx"
	"github.com/ak-the-noob-dev/authbox/models"
	"github.com/ak-the-noob-dev/authbox/utils"
	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(app *ctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input RegisterInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existing models.User
		if err := app.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}

		hash, _ := utils.HashPassword(input.Password)
		user := models.User{
			Email:        input.Email,
			PasswordHash: hash,
			Role:         "user",
		}

		if err := app.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
	}
}

func Login(app *ctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input LoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		if err := app.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		accessToken, _ := utils.CreateToken(app.JWTSecret, user.ID, user.Role, app.AccessTokenTTL, utils.AccessToken)
		refreshToken, _ := utils.CreateToken(app.JWTSecret, user.ID, user.Role, app.RefreshTokenTTL, utils.RefreshToken)

		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}
