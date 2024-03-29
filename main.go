package main

import (
	"fmt"
	"net/http"

	"github.com/ank809/SignupAPI-Golang/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", Home)
	router.GET("/signup", controllers.SignupUser)
	router.GET("/login", controllers.Loginuser)
	router.Run(":8081")

}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, "Signup to create account")
	fmt.Println("Create your account")
}
