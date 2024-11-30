package controllers

import (
	"net/http"
	"order-management/database"
	"order-management/models"
	"order-management/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Login(c *gin.Context) {

	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErrors := make(map[string]string)

		if fieldErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range fieldErrors {
				fieldName := fieldError.Field()
				tag := fieldError.Tag()
				message := formatValidationMessage(fieldName, tag)
				validationErrors[fieldName] = message
			}
		} else {
			validationErrors["error"] = err.Error()
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation error",
			"type":    "error",
			"code":    400,
			"errors":  validationErrors,
		})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ? AND password = ?", input.Username, input.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The user credentials were incorrect.",
			"type":    "error",
			"code":    400,
		})
		return
	}

	accessToken, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
			"type":    "error",
			"code":    500,
		})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate refresh token",
			"type":    "error",
			"code":    500,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token_type":    "Bearer",
		"expires_in":    432000,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func formatValidationMessage(fieldName, tag string) string {
	switch tag {
	case "required":
		return fieldName + " is required."
	default:
		return fieldName + " is invalid."
	}
}
