package main

import (
	user_controller "hotel-booking-golang-gin/controllers/user"
	"hotel-booking-golang-gin/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvironmentVariables()
}

func main() {
	r := gin.Default()

	r.POST("/ping", user_controller.CreateUser)
	r.Run() // listen and serve on 0.0.0.0:8080
}
