package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"touchedFlowed/features/user/usecases"
	"touchedFlowed/infrastructures/database"
	"touchedFlowed/infrastructures/repositories/security"
	"touchedFlowed/infrastructures/repositories/token"
	"touchedFlowed/infrastructures/repositories/user"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]
		if headerToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is empty or invalid"})
		}
		db := database.NewPgDatabase()
		userRepository := user.NewPgUserRepository(db)
		hash := security.NewHashRepository()
		tokenRepository := token.NewPgTokenRepository(database.NewRedisDatabase(), userRepository, db, hash)
		response, err := usecases.NewGetUserUseCase(&tokenRepository, &userRepository).Execute(headerToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.Set("user", response)
		c.Next()
	}
}
