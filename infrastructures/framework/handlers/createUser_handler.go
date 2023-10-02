package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"touchedFlowed/features/user/requests"
	"touchedFlowed/features/user/usecases"
	"touchedFlowed/features/utils"
	"touchedFlowed/infrastructures/repositories/database"
	"touchedFlowed/infrastructures/repositories/user"
)

func CreateUser(c *gin.Context) {
	var newUser requests.CreateUserRequest

	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is empty"})
		return
	}

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := usecases.NewCreateUserUseCase(database.NewPgUserRepository(utils.NewDatabase()), user.NewBcryptPasswordHashes()).Execute(&newUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response)
}
