package main

import (
	user_controller "hotel-booking-golang-gin/controllers/user"
	"hotel-booking-golang-gin/initializers"
	user_middlewares "hotel-booking-golang-gin/middleware/user"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.InitDatabase()
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	{
		v1 := r.Group("/api/v1")

		v1_user := v1.Group("/user")

		v1_user.POST("/create", user_middlewares.UserValidator, user_controller.CreateUser)
		v1_user.POST("/login", user_middlewares.UserLoginValidator, user_controller.LoginUser)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
