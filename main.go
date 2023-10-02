package main

import "github.com/gin-gonic/gin"
import "awesomeProject/infrastructures/framework/handlers"
import "github.com/joho/godotenv"

func main() {
	err1 := godotenv.Load()
	if err1 != nil {
		return
	}
	r := gin.Default()
	r.POST("/users", handlers.CreateUser)
	err := r.Run()
	if err != nil {
		return
	}
}
