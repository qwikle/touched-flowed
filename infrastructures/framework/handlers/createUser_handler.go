package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"touchedFlowed/features/user"
	"touchedFlowed/features/utils"
)

func CreateUser(c *gin.Context) {
	var newUser user.CreateUserRequest

	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is empty"})
		return
	}

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := user.NewCreateUserUseCase(user.NewRepository(utils.NewDatabase())).Execute(&newUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response)
}
