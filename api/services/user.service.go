package services

import (
	"api/helpers"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllUsers(c *gin.Context) {
	var users []models.User

	if err := models.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": helpers.USER_FAILED_FETCH})
		return
	}

	c.JSON(http.StatusOK, users)
}

func FindOneUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": helpers.USER_NOT_FOUND})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var createUser models.User

	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helpers.USER_FAILED_CREATE, "error": err.Error()})
		return
	}

	if err := helpers.Validation(&createUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helpers.USER_FAILED_CREATE, "error": err})
		return
	}

	var existingUser models.User
	if err := models.DB.First(&existingUser, "username = ?", createUser.Username).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": helpers.USER_ALREADY_EXIST})
		return
	}

	if err := models.DB.Create(&createUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": helpers.USER_FAILED_CREATE, "error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var updateUser models.User

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helpers.USER_FAILED_CREATE, "error": err.Error()})
		return
	}

	if err := helpers.Validation(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helpers.USER_FAILED_CREATE, "error": err})
		return
	}

	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": helpers.USER_NOT_FOUND})
		return
	}

	if err := models.DB.Model(&user).Updates(&updateUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": helpers.USER_FAILED_CREATE, "error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func UpdatePasswordUser(c *gin.Context) {
	id := c.Param("id")

	var updatePasswordUser models.UserPassword
	if err := c.ShouldBindJSON(&updatePasswordUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helpers.USER_FAILED_CREATE, "error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": helpers.USER_NOT_FOUND})
		return
	}

	if err := models.DB.Model(&user).Update("password", updatePasswordUser.Password).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": helpers.USER_FAILED_CREATE, "error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func RemoveUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": helpers.USER_NOT_FOUND})
		return
	}

	if err := models.DB.Delete(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
