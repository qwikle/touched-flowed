package main

import (
	"github.com/gin-gonic/gin"
	"touchedFlowed/infrastructures"
	"touchedFlowed/infrastructures/framework/middlewares"
)
import "touchedFlowed/infrastructures/framework/handlers"
import "github.com/joho/godotenv"

func main() {
	err1 := godotenv.Load()
	if err1 != nil {
		return
	}
	infrastructures.Init()
	r := gin.Default()
	r.POST("/sign-up", handlers.CreateUser)
	r.POST("/sign-in", handlers.SignInUser)
	r.GET("/ws", middlewares.AuthMiddleware(), handlers.UpgradeToWS)
	//todo: Finish sign-out
	//r.DELETE("/sign-out", middlewares.AuthMiddleware(), handlers.SignOutUser)
	err := r.Run()
	if err != nil {
		return
	}
}
