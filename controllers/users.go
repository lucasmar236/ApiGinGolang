package controllers

import (
	"ApiRest/models"
	"ApiRest/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	var (
		user  models.User
		param string
		id    int
	)
	param = c.Param("user")
	id, _ = strconv.Atoi(param)
	services.DB.Where("id = ?", id).
		Or("usuario LIKE ?", param).
		First(&user)
	if user.ID != 0 {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Usuário não encontrado"})
	}
}

func GetUsers(c *gin.Context) {
	var users []models.User
	services.DB.Find(&users)
	if len(users) > 0 {
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Nenhum usuário encontrado"})
	}
}

func PostUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Ocorreu um erro: " + err.Error()})
		return
	} else {
		newPass, err := bcrypt.GenerateFromPassword([]byte(user.Senha), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Senha inválida"})
		}
		user.Senha = string(newPass)
		services.DB.Create(&user)
		c.JSON(http.StatusOK, user)
		return
	}
}

func PutUser(c *gin.Context) {
	var (
		user  models.User
		param string
		id    int
	)

	param = c.Param("user")
	id, _ = strconv.Atoi(param)
	user.ID = uint(id)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Ocorreu um erro: " + err.Error()})
		return
	}
	if user.Senha != "" {
		newPass, err := bcrypt.GenerateFromPassword([]byte(user.Senha), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Senha inválida"})
			return
		}
		user.Senha = string(newPass)
	}

	services.DB.Where("id = ?", id).
		Or("usuario LIKE ?", param).
		UpdateColumns(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Atualizado com sucesso"})

}
