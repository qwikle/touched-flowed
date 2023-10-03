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

func SignInUser(c *gin.Context) {
	var newUser requests.SignInRequest

	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is empty"})
		return
	}

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := usecases.NewSignInUseCase(database.NewPgUserRepository(utils.NewDatabase()), user.NewBcryptPasswordHashes()).Execute(&newUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response)
}
