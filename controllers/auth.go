package controllers

import (
	"ApiRest/models"
	"ApiRest/services"
	"ApiRest/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c *gin.Context) {
	var (
		login models.Login
		user  models.User
	)
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	services.DB.Where("usuario LIKE ?", login.Usuario).
		First(&user)
	if login.Usuario == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Senha ou usuário inválida"})
		return
	}
	errB := bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(login.Senha))
	if errB != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Senha ou usuário inválida"})
		return
	}

	token, errJwt := utils.GenerateJwt(login.Usuario)
	if errJwt != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errJwt.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func Auth(c *gin.Context) {
	c.JSON(http.StatusNoContent, "")
}
